package basey

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Basey struct {
	db *sql.DB
}

// Make these env variables
const (
	host     = "localhost"
	port     = "5432"
	user     = "mehuljoshi"
	password = "**********"
	dbname   = "documents"
)

// Opens a new connection to the database
func OpenDatabase() *Basey {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil
	}

	fmt.Println("Connected Successfully to basey :)")
	res := new(Basey)
	res.db = db
	return res
}

func (b *Basey) InsertDocument(db *sql.DB, doc *Document) bool {
	insertSQL := `INSERT INTO documents("identifer", "title", "body") values($1, $2, $3)`
	_, e := db.Exec(insertSQL, doc.Identifier, doc.Title, doc.Body)
	if e != nil {
		return false
	}

	return true
}

/*
// // Given a phone number we look up the id of that number in the db and return
// // the id. If that number doesn't exist we return -1.
// func LookupUserId(db *sql.DB, s string) int {
// 	var id int
// 	err := db.QueryRow("SELECT id FROM USERS WHERE number like $1", s).Scan(&id)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("NO ROWS FOUND")
// 			return -1
// 		}
// 		panic(err)
// 	}
// 	return id
// }
*/

func GetDocumentById(doc_id int) *Document {
	// sql_statement := `SELECT * FROM documents WHERE id=$1`
	return nil
}
