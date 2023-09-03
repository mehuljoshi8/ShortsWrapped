package main

import (
    "recipeBot/basey"
    "net/http"
    "io"
    "github.com/gin-gonic/gin"
    "github.com/twilio/twilio-go/twiml"
    "net/url"
    "database/sql"
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


func routeInput(input string, userNumber string) string {
    id := basey.LookupUserId(db, userNumber)
    if id == -1 {
        basey.InsertUser(db, userNumber)
        id = basey.LookupUserId(db, userNumber)
    }

    if isInstaReel(input) {
        basey.InsertLink(db, id, input)
        return "inserted link :)"
    }
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
