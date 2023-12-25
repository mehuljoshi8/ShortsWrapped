package basey

import (
	"os"
	"testing"
)

// manually change these paramters to make the tests pass
// use values in .env file
// go runs tests in a different path than what the source code is in.
func setEnvs() {
	os.Setenv("PSQL_HOST", "*******")
	os.Setenv("PSQL_PORT", "*******")
	os.Setenv("PSQL_USER", "*******")
	os.Setenv("PSQL_PSWD", "*******")
	os.Setenv("PSQL_DBNAME", "*****")
}

func TestOpenDatabase(t *testing.T) {
	setEnvs()
	b := OpenDatabase()
	if b == nil {
		t.Errorf("There was an error in opening the database")
	}
	b.Close()
}

func TestInsert(t *testing.T) {
	setEnvs()
	b := OpenDatabase()
	doc := new(Document)
	doc.Identifier = "AMOLIAMOLI12"
	doc.Title = "DjungleSkog song"
	doc.Body = "amoli is a skog\nshe is a djungleskog\ndjungle djungle djungle skog skog skog"
	if done, _ := b.InsertDocument(doc); done {
		t.Errorf("Insert failed!")
	}
	b.Close()
}

func TestGetDocById(t *testing.T) {
	setEnvs()
	b := OpenDatabase()
	doc := b.GetDocumentById(1)
	if doc == nil {
		t.Errorf("Get Doc By Id Failed!")
	}
	b.Close()
}
