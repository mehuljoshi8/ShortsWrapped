package indexer

import (
	"testing"
)

func TestIndexer(t *testing.T) {
	var doc *Document = new(Document)
	doc.Id = 1
	doc.Identifier = "OCEAN78-RDum"
	doc.Title = "Craking the Coding Interview"
	doc.Body = "The code below offers an inital solution, but it has a bug in it. Can you find it"
	doc.Body += "\nThe problem with this code occurs in the case where\nSuppose we call commonAncestor(node 3, node 5, node 7)."
	var i *Indexer = NewIndexer()
	i.Index(doc)
	var doc2 *Document = new(Document)
	doc2.Id = 2
	doc2.Identifier = "CHILLIFIRE12"
	doc2.Title = "Sriracha Ingredients"
	doc2.Body = "Chilli, Water, Sugar, Distilled Vinegar, Garlic, Salt, Thickener: Modified tapicao starch. Xanthan Gum"
	i.Index(doc2)

}
