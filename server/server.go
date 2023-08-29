package server

import (
    "fmt"
    "net/http"
    "io"
    "github.com/gin-gonic/gin"
    "github.com/twilio/twilio-go/twiml"
    "net/url"
)

func smsHandler(context *gin.Context) {
    body, err := io.ReadAll(context.Request.Body)
    if err != nil {
        context.Abort()
        return
    }

    requestParams, err := url.ParseQuery(string(body))
    if err != nil {
        context.Abort()
        return
    }

    fmt.Println(requestParams["Body"])
 
    for k := range requestParams {
        fmt.Println(k)
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

func PrintHello() {
    fmt.Println("Hello Modules! This is the server speaking!")
}
