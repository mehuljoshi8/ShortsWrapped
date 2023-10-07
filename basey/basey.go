package basey

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Opens a new connection to the database
func OpenDatabase() *sql.DB {
	host := os.Getenv("PSQL_HOST")
	port := os.Getenv("PSQL_PORT")
	user := os.Getenv("PSQL_USER")
	dbname := os.Getenv("RBOT_DBNAME")
	password := os.Getenv("PSQL_PW")
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	fmt.Println("Connected Successfully to basey :)")
	return db
}

// Inserts a new user associated with their phone number (number) into db
func InsertUser(db *sql.DB, number string) {
	insertSQL := `INSERT INTO "users"("number") values($1)`
	_, e := db.Exec(insertSQL, number)
	checkError(e)
}

// Inserts a Link associated with a user_id and a reel identifer into the db
func InsertLink(db *sql.DB, userId int, link string) {
	insertSQL := `INSERT INTO "links"("user_id", "hyperlink") values($1, $2)`
	_, e := db.Exec(insertSQL, userId, link)
	checkError(e)
}

// Checks if there is an error and panics
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Given a phone number we look up the id of that number in the db and return
// the id. If that number doesn't exist we return -1.
func LookupUserId(db *sql.DB, s string) int {
	var id int
	err := db.QueryRow("SELECT id FROM USERS WHERE number like $1", s).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("NO ROWS FOUND")
			return -1
		}
		panic(err)
	}
	return id
}

// Returns the links for a given user with user.id.
func GetLinksForUser(db *sql.DB, user_id int) ([]Link, error) {
	var links []Link
	rows, err := db.Query("SELECT * FROM LINKS WHERE user_id=$1", user_id)
	if err != nil {
		return nil, fmt.Errorf("GetLinksForUser (%v)", user_id)
	}

	defer rows.Close()
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.Id, &link.Identifer, &link.UserId)
		if err != nil {
			return links, err
		}
		links = append(links, link)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return links, nil
}
