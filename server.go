package main

import (
    "fmt"
    "net/http"
    "io"
    "github.com/gin-gonic/gin"
    "github.com/twilio/twilio-go/twiml"
    "net/url"
)

// TODO: Complete database work
// users, links, etc, etc tables
// pulling, querying, etc, etc, etc
// then link that to this

// also implement a tinyURL server so that
// links can live on the server here instead of elsewhere

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
    router := gin.Default()
    router.POST("/sms", smsHandler)
    router.Run(":4040")
}
