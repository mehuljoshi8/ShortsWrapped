package indexer

import (
	"fmt"
	"recipeBot/basey"
	"strings"
	"unicode"
)

// Type alias for basey.Document becuase that struct is used heavily through this API
type Document basey.Document

// The indexer is what most of the code in this file is concerend about.
type Indexer struct {
	doc_set map[uint64]struct{}
	index   map[string]map[uint64][]int
}

// Returns a new instance of an Indexer
func NewIndexer() *Indexer {
	i := new(Indexer)
	i.doc_set = make(map[uint64]struct{})
	i.index = make(map[string]map[uint64][]int)
	return i
}

// The index function puts the document into the index if
// the index already contains that document then we don't
// index it and return false. Otherwise we index it and return true.
func (i *Indexer) Index(doc *Document) bool {
	if _, found := i.doc_set[doc.Id]; found {
		return false
	}

	i.doc_set[doc.Id] = struct{}{}
	tokens := tokenize(doc.Body)
	fmt.Println(tokens)

	return true
}

// Tokenize a string into an array of strings
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// The single input to process a query; returns a list of doc_ids sorted
// based on the query that we are given.
func ProcessQuery(input string) []uint64 {
	return make([]uint64, 0)
}
