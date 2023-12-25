package indexer

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	var s string = "The next day - Sunday - Oren drove me to ortega, McNarma, and Jones to see the Red Will"
	fmt.Println(len(tokenize(s)))
}
