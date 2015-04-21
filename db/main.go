package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(db *sql.DB) (sql.Result, error) {
	query := `
	CREATE TABLE user (
		id				INTEGER NOT NULL PRIMARY KEY,
		username		TEXT
	)
	`

	return db.Exec(query)
}

func generateData(db *sql.DB) (sql.Result, error) {
	query := `
	INSERT INTO user (username) VALUES ('Smith')
	`
	return db.Exec(query)
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = initDB(db)
	checkDBExecutionError(err)

	result, err := generateData(db)
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

func checkDBExecutionError(err error) {
	if err != nil {
		log.Printf("cannot execute query: %q\n", err)
		return
	}
}
