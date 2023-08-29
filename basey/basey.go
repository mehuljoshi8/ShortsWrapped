package basey

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "os"
)

const (
    host = "localhost"
    port = 5432
    user = "postgres"
    dbname = "recipes"
)

var password = os.Getenv("PSQL_PW")

/*
CREATE TABLE USERS (
    id SERIAL PRIMARY KEY,
    number varchar(12) UNIQUE NON NULL
);
*/
type User struct {
    id          uint64
    number      string
}

/*
CREATE TABLE LINKS (
    id SERIAL PRIMARY KEY,
    hyperlink varchar NON NULL,
    user_id INTEGER REFERENCES users (id)
);
*/
type Link struct {
    id          uint64
    userid      uint64
    hyperlink   string
}

func OpenDatabase() *sql.DB {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlconn)
    checkError(err)
    fmt.Println("Connected Successfully to basey :)")
    //FindUser(db, "+14259431674")
    //InsertLink(db, 1, "https://www.instagram.com/reel/Cv49nyaLZpx/?utm_source=ig_web_copy_link&igshid=MzRlODBiNWFlZA==")
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
    insertSQL := `INSERT INTO "links"("user_id", "hyperlink") values($1, $2)`
    _, e := db.Exec(insertSQL, userId, link)
    checkError(e)
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

// Returns the id for a number in the users database
func LookupUserId(db *sql.DB, s string) int {
    fmt.Println(s)
    var id int
    err := db.QueryRow("SELECT id FROM USERS WHERE number like $1", s).Scan(&id)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("NO ROWS FOUND")
            return -1;
        }
        panic(err)
    }
    fmt.Print("id = ")
    fmt.Println(id)
    return id
}
