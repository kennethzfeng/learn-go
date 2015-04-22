package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/codegangsta/cli"
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
	result, err := db.Exec(query, "Smith")
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

func selectData(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user")
	checkDBExecutionError(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var username string

		if err := rows.Scan(&id, &username); err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d Username: %s\n", id, username)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := cli.NewApp()
	app.Name = "Database Demo Utility"
	app.Version = "0.1"
	app.Author = "Kenneth Feng"

	app.Commands = []cli.Command{
		{
			Name: "init",
			Action: func(c *cli.Context) {
				initDB(db)
			},
		},
		{
			Name: "insert",
			Action: func(c *cli.Context) {
				generateData(db)
			},
		},
		{
			Name: "select",
			Action: func(c *cli.Context) {
				selectData(db)
			},
		},
	}

	app.Run(os.Args)
}

func checkDBExecutionError(err error) {
	if err != nil {
		log.Printf("cannot execute query: %q\n", err)
		return
	}
}
