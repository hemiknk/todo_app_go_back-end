package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("can't load .env file: %w", err)
	}

	env := os.Getenv("DEV_ENV")

	var db *sql.DB
	var driverName string

	switch env {
	case "true":
		dbPath := os.Getenv("SQLITE_DB_PATH")
		driverName = "sqlite3"
		var err error

		db, err = sql.Open(driverName, dbPath)
		if err != nil {
			return nil, err
		}
	case "false":
		host := os.Getenv("PG_DB_HOST")
		port := os.Getenv("PG_DB_PORT")
		user := os.Getenv("PG_DB_USER")
		password := os.Getenv("PG_DB_PASSWORD")
		dbName := os.Getenv("PG_DB_NAME")
		driverName = "postgres"
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
		var err error

		db, err = sql.Open(driverName, connStr)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown env var DEV_ENV expected true or false")
	}

	return db, nil
}
