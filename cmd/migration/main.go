package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/hemiknk/todo_app_go_back-end/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		panic("Please provide an argument: up or down")
	}

	if os.Args[1] == "up" {
		err = MigrateUp(conn)
		if err != nil {
			panic(err)
		}
		log.Println("Migration up was successful")

		return
	}

	if os.Args[1] == "down" {
		err = MigrateDown(conn)
		if err != nil {
			panic(err)
		}
		log.Println("Migration down was successful")

		return
	}

	log.Println("unknown command, expected up or down")
}

func MigrateUp(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS todo (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			title TEXT,
			done BOOLEAN,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP
		);
	`

	log.Println("Migrationg up...")

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func MigrateDown(db *sql.DB) error {
	query := `
		DROP TABLE IF EXISTS todo;
	`
	log.Println("Migrationg down...")

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
