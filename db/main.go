package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := `
	CREATE TABLE user (
		id				INTEGER NOT NULL PRIMARY KEY,
		username		TEXT
	)
	`

	_, err = db.Exec(query)
	if err != nil {
		log.Printf("cannot execute query:'\n%s\n' due to %q", query, err)
		return
	}

}
