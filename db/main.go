package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(db *sql.DB) {
	query := `
	CREATE TABLE user (
		id				INTEGER NOT NULL PRIMARY KEY,
		username		TEXT
	)
	`

	_, err := db.Exec(query)
	checkDBExecutionError(err)
}

func generateData(db *sql.DB) {
	query := `
	INSERT INTO user (username) VALUES (?)
	`
	result, err := db.Exec(query)
	checkDBExecutionError(err)

	insertID, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
	}

	log.Printf("LastInsertId(): %d", insertID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	log.Printf("RowsAffected(): %d", rowsAffected)
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	initDB(db)

	generateData(db)
}

func checkDBExecutionError(err error) {
	if err != nil {
		log.Printf("cannot execute query: %q\n", err)
		return
	}
}
