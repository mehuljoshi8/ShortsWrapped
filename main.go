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
)

var db *sql.DB

// to select instareel caption
// see document.querySelector("._a9zs");
func isInstaReel(s string) bool {
    const instaReelStarter = "https://www.instagram.com/reel/"
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

// returns a string result that is outputted to the user
// based on an interaction by the input.
func routeInput(input string, userNumber string) string {
    userid := basey.LookupUserId(db, userNumber)
    if userid == -1 {
        basey.InsertUser(db, userNumber)
        userid = basey.LookupUserId(db, userNumber)
    }

    if isInstaReel(input) {
        basey.InsertLink(db, userid, input)
        return "inserted link :)"
    } 
    // we are going to make an assumption that
    // if the user doesn't type a link then we have
    // to make sense of what the user wants
    // we are going to do this the super cool way
    // the 333 way parallel search systems way. 
    fmt.Println("input = ", input)

    // step 1 extract the recipes from the links
    links, err := basey.GetLinksForUser(db, userid)
    if err != nil {
       return "error"
    }
    fmt.Println(len(links))
    for i := 0; i < len(links); i++ {
        fmt.Println(links[i].Hyperlink)
    }
    // now we have all the links
    // now the next step is to extract the content out
    // the links and store it in memory in some kinda a 
    // data structure prolly linkid -> string
    // then index that datastructure to search 333 way...
    // this project is kinda fun ngl
    return "working on other features" 
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
