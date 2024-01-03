package indexer

import (
	"recipeBot/basey"
	"strings"
	"unicode"

	snowballeng "github.com/kljensen/snowball/english"
)

// The indexer is what most of the code in this file is concerend about.
type Indexer struct {
	doc_set map[uint64]struct{}
	index   map[string]map[uint64][]uint64
}

// All The Stop Words that we use for the indexer
var stopWords = map[string]struct{}{
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

// Returns a new instance of an Indexer
func NewIndexer() *Indexer {
	i := new(Indexer)
	i.doc_set = make(map[uint64]struct{})
	i.index = make(map[string]map[uint64][]uint64)
	return i
}

// given a string token returns a token that is the result of applying
// the lowercase, stopword, and stem filters on it.
func transformToken(token string) string {
	if len(token) > 0 {
		token = strings.ToLower(token)
		if _, ok := stopWords[token]; !ok {
			token = snowballeng.Stem(token, false)
			return token
		}
	}
	return ""
}

// Appends a position onto the index for a given docId and token
func (i *Indexer) appendPos(token string, docId uint64, startPos uint64) {
	if len(i.index[token]) == 0 {
		i.index[token] = make(map[uint64][]uint64)
	}
	i.index[token][docId] = append(i.index[token][docId], startPos)
}

// The index function puts the document into the index if
// the index already contains that document then we don't
// index it and return false. Otherwise we index it and return true.
func (i *Indexer) Index(doc *basey.Document) bool {
	if _, found := i.doc_set[doc.Id]; found {
		return false
	}

	i.doc_set[doc.Id] = struct{}{}

	var token string
	var token_builder strings.Builder
	var startPos uint64 = 0
	for j, c := range doc.Body {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			token = transformToken(token_builder.String())
			if len(token) > 0 {
				i.appendPos(token, doc.Id, startPos)
			}
			startPos = uint64(j) + 1
			token_builder.Reset()
		} else {
			token_builder.WriteRune(c)
		}
	}

	token = transformToken(token_builder.String())
	if len(token) > 0 {
		i.appendPos(token, doc.Id, startPos)
	}

	return true
}

// The single input to process a query; returns a list of doc_ids sorted
// based on the query that we are given.
// The rank function of a document is what we have to figure out to make
// our search engine a good one.
func (i *Indexer) ProcessQuery(query string) map[uint64]uint64 {
	var results map[uint64]uint64 = make(map[uint64]uint64, 0)

	var token string
	var token_builder strings.Builder = strings.Builder{}
	var tokenCount uint64 = 0
	// var startPos uint64 = 0
	for _, c := range query {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			token = transformToken(token_builder.String())
			if len(token) > 0 {
				for doc_id, lst := range i.index[token] {
					results[doc_id] += uint64(len(lst))
				}
			}
			// startPos = uint64(j) + 1
			token_builder.Reset()
		} else {
			token_builder.WriteRune(c)
		}
	}

	// for the last token
	token = transformToken(token_builder.String())
	if len(token) > 0 {
		tokenCount++
		for doc_id, lst := range i.index[token] {
			results[doc_id] += uint64(len(lst))
		}
	}

	return results
}
