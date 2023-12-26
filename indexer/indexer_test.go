package indexer

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	var s string = "The next day - Sunday - Oren drove me to ortega, McNarma, and Jones to see the Red Will"
	fmt.Println(len(tokenize(s)))
}

func TestIndexer(t *testing.T) {
	var doc *Document = new(Document)
	doc.Id = 1
	doc.Identifier = "OCEAN78-RDum"
	doc.Title = "Craking the Coding Interview"
	doc.Body = "The code below offers an inital solution, but it has a bug in it. Can you find it"
	doc.Body += "\nThe problem with this code occurs in the case where\nSuppose we call commonAncestor(node 3, node 5, node 7)."
	var i *Indexer = NewIndexer()
	i.Index(doc)

}
