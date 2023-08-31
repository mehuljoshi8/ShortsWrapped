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
func isInstaReelLink(s string) bool {
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

// TODO: implement a tiny URL server
// where we can shorten links and store them.
func smsHandler(context *gin.Context) {
    body, err := io.ReadAll(context.Request.Body)
    if err != nil {
        context.Abort()
        return 
    }

    // type(reqParams) = map[string]string
    reqParams, err := url.ParseQuery(string(body))
    if err != nil {
        context.Abort()
        return
    }
    
    input := reqParams["Body"][0]
    userNumber := reqParams["From"][0]
    id := basey.LookupUserId(db, userNumber)
    if id == -1 {
        basey.InsertUser(db, userNumber)
        id = basey.LookupUserId(db, userNumber)
    }

    if isInstaReelLink(input) {
        basey.InsertLink(db, id, input)
    }

    message := &twiml.MessagingMessage {
        Body: "Let the chaos begin",
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
