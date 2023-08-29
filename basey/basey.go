package basey

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)
const (
    host = "local"
    port = 5432
    user = "mehuljoshi"
    password = "46AGMDJS"
    dbname = "recipies"
)

type User struct {
    id          uint64
    number      string
}

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


    s := "+14259431672"
    var id uint64
    err = db.QueryRow("SELECT FROM USERS WHERE number=?", s).Scan(&id)
    checkError(err)
    fmt.Print("id = ")
    fmt.Println(id)

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

func LookupUser(number string) {
    //var id uint64
    //err := db.QueryRow("SELECT FROM USERS WHERE number = ?", number).Scan(&id)
    //if err != nil {
    //    if err == sql.ErrNoRows {
    //        fmt.Println("This number is not in the database")
    //    }
    //    checkError(err)
    //}
    //fmt.Println(id)
    fmt.Println(number)
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func FindUser(db *sql.DB, s string) {
    fmt.Println(s)
    var id uint64
    err := db.QueryRow("SELECT FROM USERS WHERE number=?", s).Scan(&id)
    checkError(err)
    fmt.Print("id = ")
    fmt.Println(id)
}
