package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hemiknk/todo_app_go_back-end/internal/db"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("can't load .env file: %v", err))
	}

	err = db.SetUpConnection()
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		panic("Please provide an argument: up or down")
	}

	if os.Args[1] == "up" {
		err = MigrateUp()
		if err != nil {
			panic(err)
		}
		log.Println("Migration up was successful")

		return
	}

	if os.Args[1] == "down" {
		err = MigrateDown()
		if err != nil {
			panic(err)
		}
		log.Println("Migration down was successful")

		return
	}

	log.Println("unknown command, expected up or down")
}

func MigrateUp() error {
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

	_, err := db.Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func MigrateDown() error {
	query := `
		DROP TABLE IF EXISTS todo;
	`
	log.Println("Migrationg down...")

	_, err := db.Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
