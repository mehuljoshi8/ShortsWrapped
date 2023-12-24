package basey

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Wrapper struct to not expose all the functionality of the sql.DB to clients of basey
type Basey struct {
	db *sql.DB
}

// Opens a new connection to the database
func OpenDatabase() *Basey {
	godotenv.Load(".env")
	host := os.Getenv("PSQL_HOST")
	port := os.Getenv("PSQL_PORT")
	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PSWD")
	dbname := os.Getenv("PSQL_DBNAME")

	psqlconn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil
	}
	res := new(Basey)
	res.db = db
	return res
}

// Inserts a document into the database
func (b *Basey) InsertDocument(doc *Document) bool {
	insertSQL := `INSERT INTO documents("identifer", "title", "body") values($1, $2, $3)`
	_, err := b.db.Exec(insertSQL, doc.Identifier, doc.Title, doc.Body)
	return err == nil
}

// Closes the connection to the database
func (b *Basey) Close() {
	b.db.Close()
}

// Returns a document associated by the doc_id if it exists in the database
func (b *Basey) GetDocumentById(doc_id uint64) *Document {
	sql_statement := `SELECT * FROM documents WHERE id=$1`
	var doc *Document = new(Document)
	doc.Id = doc_id
	err := b.db.QueryRow(sql_statement, doc_id).Scan(&doc.Id, &doc.Identifier, &doc.Title, &doc.Body)
	if err != nil {
		return nil
	}
	return doc
}
