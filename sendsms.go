package main

import (
    "fmt"
    "github.com/twilio/twilio-go"
    api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
    client := twilio.NewRestClient()

    params := &api.CreateMessageParams{}
    params.SetBody("Go Bot")
    params.SetFrom("+1*********")
    params.SetTo("+14259431672")
    resp, err := client.Api.CreateMessage(params)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(resp.Sid)
    }
}
