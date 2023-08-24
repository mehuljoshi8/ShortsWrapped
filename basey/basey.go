package basey

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "mehuljoshi"
    password = "46AGMDJS"
    dbname = "recipies"
)

// Returns an open connection to our psql db
func OpenDatabase() (*sql.DB) {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlconn)
    checkError(err)
    return db
}

// inserts numbers into db
func InsertUser(db *sql.DB, number string) {
    insertSQL := `INSERT INTO "users"("number") values($1)`
    _, e := db.Exec(insertSQL, number)
    checkError(e)
}

// insert links into the db
func InsertLink(db *sql.DB, userId int, link string) {
    insertSQL := `INSERT INTO "links"("userid", "hyperlink") values($1, $2)`
    _, e := db.Exec(insertSQL, userId, link)
    checkError(e)
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func PrintString(s string) {
    fmt.Println(s)
}

