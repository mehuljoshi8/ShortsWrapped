package main

// var db *sql.DB

// // ============== Building a full-text search engine ==============
// // we are going to be using an inverted index.... :)
// func tokenize(text string) []string {
// 	return strings.FieldsFunc(text, func(r rune) bool {
// 		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
// 	})
// }

// func lowercaseFilter(tokens []string) []string {
// 	r := make([]string, len(tokens))
// 	for i, token := range tokens {
// 		r[i] = strings.ToLower(token)
// 	}
// 	return r
// }

// func stopwordFilter(tokens []string) []string {
// 	var stopWords = map[string]struct{}{
// 		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
// 		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
// 	}

// 	r := make([]string, 0, len(tokens))
// 	for _, token := range tokens {
// 		if _, ok := stopWords[token]; !ok {
// 			r = append(r, token)
// 		}
// 	}
// 	return r
// }

// func stemmerFilter(tokens []string) []string {
// 	r := make([]string, len(tokens))
// 	for i, token := range tokens {
// 		r[i] = snowballeng.Stem(token, false)
// 	}
// 	return r
// }

// func analyze(text string) []string {
// 	tokens := tokenize(text)
// 	tokens = lowercaseFilter(tokens)
// 	tokens = stopwordFilter(tokens)
// 	tokens = stemmerFilter(tokens)
// 	return tokens
// }

// func intersection(a []int, b []int) []int {
// 	maxLen := len(a)
// 	if len(b) > maxLen {
// 		maxLen = len(b)
// 	}
// 	r := make([]int, 0, maxLen)
// 	var i, j int
// 	for i < len(a) && j < len(b) {
// 		if a[i] < b[j] {
// 			i++
// 		} else if a[i] > b[j] {
// 			j++
// 		} else {
// 			r = append(r, a[i])
// 			i++
// 			j++
// 		}
// 	}
// 	return r
// }

// // =============== End of Search Engine ========================

// // returns a string result that is outputted to the user
// // based on an interaction by the input.
// func routeInput(input string, userNumber string) string {
// 	userid := getUserId(userNumber)

// 	if isInstaReel(input) {
// 		basey.InsertLink(db, userid, getReelIdentifer(input))
// 		return "inserted reel identifer"
// 	}
// 	// fmt.Println("input = " + input)
// 	// links, err := basey.GetLinksForUser(db, userid)
// 	// if err != nil {
// 	// 	return "error"
// 	// }

// 	// // TODO: move recipes to it's own table.
// 	// recipes := make([]string, len(links))
// 	// for i, link := range links {
// 	// 	recipes[i] = scrapeRecipe(link.ReelIdentifer)
// 	// }

// 	// var index map[string][]int = make(map[string][]int)
// 	// // build our index here
// 	// for i, r := range recipes {
// 	// 	for _, token := range analyze(r) {
// 	// 		ids := index[token]
// 	// 		if ids != nil && ids[len(ids)-1] == i {
// 	// 			continue
// 	// 		}
// 	// 		index[token] = append(ids, i)
// 	// 	}
// 	// }

// 	// var r []int
// 	// for _, token := range analyze(input) {
// 	// 	if ids, ok := index[token]; ok {
// 	// 		if r == nil {
// 	// 			r = ids
// 	// 		} else {
// 	// 			r = intersection(r, ids)
// 	// 		}
// 	// 	}
// 	// }

// 	// fmt.Println(r)
// 	// // for now i don't care lets just build the search algo
// 	// for _, j := range r {
// 	// 	fmt.Println(recipes[j])
// 	// 	fmt.Println("================================================")
// 	// }

// 	return "working on search"
// }

// // Handles the SMS input for the gin sever.
// func smsHandler(context *gin.Context) {
// 	requestParams := getRequestParameters(context)
// 	if requestParams == nil {
// 		return
// 	}

// 	res := routeInput(
// 		requestParams["Body"][0],
// 		requestParams["From"][0])

// 	fmt.Println(res)

// 	message := &twiml.MessagingMessage{
// 		Body: res,
// 	}

// 	twimlResult, err := twiml.Messages([]twiml.Element{message})
// 	if err != nil {
// 		context.String(http.StatusInternalServerError, err.Error())
// 	} else {
// 		context.Header("Content-Type", "text/xml")
// 		context.String(http.StatusOK, twimlResult)
// 	}
// }

// func main() {
// 	db = basey.OpenDatabase()
// 	defer db.Close()
// 	router := gin.Default()
// 	router.POST("/sms", smsHandler)
// 	router.Run(":4040")
// }

func main() {

}
