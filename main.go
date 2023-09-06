package main

import (
    "recipeBot/basey"
    "net/http"
    "io"
    "github.com/gin-gonic/gin"
    "github.com/twilio/twilio-go/twiml"
    "net/url"
    "database/sql"
    "fmt"
    "github.com/PuerkitoBio/goquery"
)

var db *sql.DB

const instaReelStarter = "https://www.instagram.com/reel/"

// to select instareel caption
// see document.querySelector("._a9zs");
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

// Instagram reels are in the form of
// https://www.instagram.com/reel/<IDENTIFER>/id=...
// This function just returns the IDENTIFER for the reel.
func getReelIdentifer(link string) string {
    i := len(link) - 1
    for link[i] != '/' {
        i--
    }

    fmt.Println(link[len(instaReelStarter):i])
    return link[len(instaReelStarter):i]
}

// Parses the body and returns the request parameters.
// which is nothing more than a map of string -> []string
// (denoted Values) as by the url package.
// returns nil on error.
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

// extracts the recipe for the reelId and
// outs it to the console
func scrapeRecipe(reelId string) {
    res, err := http.Get(instaReelStarter + reelId + "/")
    if err != nil {
        return
    }

    defer res.Body.Close()
    if res.StatusCode != 200 {
        return
    }

    doc, _ := goquery.NewDocumentFromReader(res.Body)
    doc.Find("meta").Each(func(i int, s *goquery.Selection) {
        if name, _ := s.Attr("property"); name == "og:title" {
            content, _ := s.Attr("content")
            // content is what the recipe is stored in if
            // it's in the little 3 bubbles
            fmt.Println(content)
        }
    })
}

// returns a string result that is outputted to the user
// based on an interaction by the input.
func routeInput(input string, userNumber string) string {
    userid := basey.LookupUserId(db, userNumber)
    if userid == -1 {
        basey.InsertUser(db, userNumber)
        userid = basey.LookupUserId(db, userNumber)
    }

    if isInstaReel(input) {
        basey.InsertLink(db, userid, getReelIdentifer(input))
        return "inserted reel identifer"
    }

    
    // working on search feature
    links, err := basey.GetLinksForUser(db, userid)
    if err != nil {
        return "error"
    }

    for _, link := range links {
        scrapeRecipe(link.ReelIdentifer)
    }

    return "working on search"
}

// TODO: implement a tiny URL server
// where we can shorten links and store them.
func smsHandler(context *gin.Context) {
    requestParams := getRequestParameters(context)
    if requestParams == nil {
        return
    }
    
    res := routeInput(
            requestParams["Body"][0],
            requestParams["From"][0])
    
    fmt.Println(res)
    message := &twiml.MessagingMessage {
        Body: res,
    }

    twimlResult, err := twiml.Messages([]twiml.Element{message})
    if err != nil {
        context.String(http.StatusInternalServerError, err.Error())
    } else {
        context.Header("Content-Type", "text/xml")
        context.String(http.StatusOK, twimlResult)
    }
}

func main() {
    db = basey.OpenDatabase()
    defer db.Close()
    router := gin.Default()
    router.POST("/sms", smsHandler)
    router.Run(":4040")
}
