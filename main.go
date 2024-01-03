package main

import (
	"net/http"
	"recipeBot/basey"
	"recipeBot/indexer"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go/twiml"
)

var db *basey.Basey
var i *indexer.Indexer

// returns a string result that is outputted to the user
// based on an interaction by the input.
func routeInput(input string, userNumber string) string {
	if isInstaReel(input) {
		var doc *basey.Document
		var err error
		doc = new(basey.Document)
		doc.Identifier = getReelIdentifier(input)
		err, doc.Body = scrapeRecipe(doc.Identifier)
		if err != nil {
			return "Insert failed. Link doesn't point to a valid instagram reel"
		}
		inserted, tmp_id := db.InsertDocument(doc)
		if !inserted {
			return "Insert failed. Input is malformed"
		}

		doc.Id = uint64(tmp_id)

		i.Index(doc)

		return "Insert Succeeded."
	}
	return "working on search"
}

// // Handles the SMS input for the gin sever.
func smsHandler(context *gin.Context) {
	requestParams := getRequestParameters(context)
	if requestParams == nil {
		return
	}

	res := routeInput(
		requestParams["Body"][0],
		requestParams["From"][0])

	message := &twiml.MessagingMessage{
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
	i = indexer.NewIndexer()
	defer db.Close()
	router := gin.Default()
	router.POST("/sms", smsHandler)
	router.Run(":4040")
}
