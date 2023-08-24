package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "mehuljoshi"
    password = ""
    dbname = "recipies"
)

func main() {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)

    defer db.Close()

    // dynamic
    insertDynStmt := `insert into "users"("id", "number") values($1, $2)`
    _, e := db.Exec(insertDynStmt, 1, "+14259431672")
    CheckError(e)
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
