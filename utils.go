package main

import (
	"io"
	"net/http"
	"net/url"
	"recipeBot/basey"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const instaReelStarter string = "https://www.instagram.com/reel/"

// The isInstaReel function sees if a given string is in the form
// of an instagram reel. Which is usually prefixed by the instaReelStarter
// if it is an instagram reel we return true otherwise we return false.
func isInstaReel(s string) bool {
	if len(s) < len(instaReelStarter) {
		return false
	}

	for i := 0; i < len(instaReelStarter); i++ {
		if instaReelStarter[i] != s[i] {
			return false
		}
	}
	return true
}

// RecipeBot currently only recives reels in the form of
// instaReelStarter + IDENTIFER + /id=....
// This function returns the IDENTIFER for the reel.
func getReelIdentifer(link string) string {
	i := len(link) - 1
	for link[i] != '/' {
		i--
	}

	return link[len(instaReelStarter):i]
}

// The getRequestParameters function takes in a *gin.Context
// (which contains all the information about the request that the
// handler might need to process it), parses
// the body and returns the request parameters
func getRequestParameters(context *gin.Context) url.Values {
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.Abort()
		return nil
	}

	requestParams, err := url.ParseQuery(string(body))
	if err != nil {
		context.Abort()
		return nil
	}

	return requestParams
}

// TODO: fix the error cases to not just return error and actually return the error.
// The scapeRecipe function takes in a reelId and attempts to scrape the content
// from the reel and returns the content through a string. If we can't scrape
// the content for some reason we return an error and an empty string.
func scrapeRecipe(reelId string) string {
	res, err := http.Get(instaReelStarter + reelId + "/")
	if err != nil {
		return "error"
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "error"
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	recipeContent := ""
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("property"); name == "og:title" {
			content, _ := s.Attr("content")
			recipeContent += content
		}
	})

	return recipeContent
}

// Looks up the userId for a number. If that number
// is not inserted in the database we insert it and then
// return the id for that newly inserted value.
func getUserId(number string) int {
	userid := basey.LookupUserId(db, number)
	if userid == -1 {
		basey.InsertUser(db, number)
		userid = basey.LookupUserId(db, number)
	}

	return userid
}
