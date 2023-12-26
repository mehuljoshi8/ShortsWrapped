package indexer

import (
	"recipeBot/basey"
	"strings"
	"unicode"

	snowballeng "github.com/kljensen/snowball/english"
)

// Type alias for basey.Document becuase that struct is used heavily through this API
type Document basey.Document

// The indexer is what most of the code in this file is concerend about.
type Indexer struct {
	doc_set map[uint64]struct{}
	index   map[string]map[uint64][]uint64
}

// Returns a new instance of an Indexer
func NewIndexer() *Indexer {
	i := new(Indexer)
	i.doc_set = make(map[uint64]struct{})
	i.index = make(map[string]map[uint64][]uint64)
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

	var stopWords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}

	var token string
	var token_builder strings.Builder
	var startPos uint64 = 0
	for j, c := range doc.Body {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			token = token_builder.String()
			if len(token) > 0 {
				token = strings.ToLower(token)
				if _, ok := stopWords[token]; !ok {
					token = snowballeng.Stem(token, false)
					if len(i.index[token]) == 0 {
						i.index[token] = make(map[uint64][]uint64)
					}
					i.index[token][doc.Id] = append(i.index[token][doc.Id], startPos)
				}
			}
			startPos = uint64(j) + 1
			token_builder = strings.Builder{}
		} else {
			token_builder.WriteRune(c)
		}
	}
	return true
}

// The single input to process a query; returns a list of doc_ids sorted
// based on the query that we are given.
func ProcessQuery(query string) []uint64 {
	return make([]uint64, 0)
}
