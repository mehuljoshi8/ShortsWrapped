package main

import (
    "recipeBot/basey"
    "fmt"
    "net/http"
    "io"
    "github.com/gin-gonic/gin"
    "github.com/twilio/twilio-go/twiml"
    "net/url"
    "database/sql"
)

// TODO: Move handler to server.go file
// Implement Server struct that is incharge of managing the db.

var db *sql.DB

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
    
    fmt.Println(reqParams["Body"])
    fmt.Println(reqParams["From"])

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
    //basey.FindUser(db, "+14259431672")
    router := gin.Default()
    router.POST("/sms", smsHandler)
    router.Run(":4040")
}
