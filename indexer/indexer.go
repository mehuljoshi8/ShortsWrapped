package indexer

import (
	"strings"
	"unicode"
)

// Tokenize a string into an array of strings
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
