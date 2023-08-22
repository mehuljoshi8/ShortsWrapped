package main

import (
    "fmt"
    "net/http"
    "io"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/twilio/twilio-go/twiml"
)

// TODO: Implement a function that takes the string
// body and returns a map from string -> string.
// stringToJsonMap()

func main() {
    router := gin.Default()

    router.POST("/sms", func(context *gin.Context) {
        body, err := io.ReadAll(context.Request.Body)
        fmt.Println("hello world")
        if err != nil {
            os.Exit(0)    
        }

        fmt.Println(string(body))
        message := &twiml.MessagingMessage {
            Body: "Amoli is coming! RUNNNN",
        }

        twimlResult, err := twiml.Messages([]twiml.Element{message})
        if err != nil {
            context.String(http.StatusInternalServerError, err.Error())
        } else {
            context.Header("Content-Type", "text/xml")
            context.String(http.StatusOK, twimlResult)
        }
    })

    router.Run(":4040")
}
