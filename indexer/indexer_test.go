package indexer

import (
	"fmt"
	"recipeBot/basey"
	"testing"
)

func TestIndexer(t *testing.T) {
	var doc *basey.Document = new(basey.Document)
	doc.Id = 1
	doc.Identifier = "OCEAN78-RDum"
	doc.Title = "Craking the Coding Interview"
	doc.Body = "The code below offers an inital solution, but it has a bug in it. Can you find it"
	doc.Body += "\nThe problem with this code occurs in the case where\nSuppose we call commonAncestor(node 3, node 5, node 7)."
	var i *Indexer = NewIndexer()
	i.Index(doc)
	var doc2 *basey.Document = new(basey.Document)
	doc2.Id = 2
	doc2.Identifier = "CHILLIFIRE12"
	doc2.Title = "Sriracha Ingredients"
	doc2.Body = "Chilli, Water, Sugar, Distilled Vinegar, Garlic, Salt, Thickener: Modified tapicao starch. Xanthan Gum"
	i.Index(doc2)
}

func TestQueryProcessor(t *testing.T) {
	var doc *basey.Document = new(basey.Document)
	doc.Id = 1
	doc.Identifier = "OCEAN78-RDum"
	doc.Title = "Craking the Coding Interview"
	doc.Body = "The code below offers an inital solution, but it has a bug in it. Can you find it"
	doc.Body += "\nThe problem with this code occurs in the case where\nSuppose we call commonAncestor(node 3, node 5, node 7)."

	var doc2 *basey.Document = new(basey.Document)
	doc2.Id = 2
	doc2.Identifier = "GRISHAM390HJ"
	doc2.Title = "The Pelican Breif"
	doc2.Body = "The next day America learns that two of its supreme court justices have been assassinated. Law student solves the problem."

	doc2.Body = "solve                                            B"
	var i *Indexer = NewIndexer()
	i.Index(doc)
	i.Index(doc2)
	fmt.Println(i.ProcessQuery("code offers a way to solve problems"))
}
