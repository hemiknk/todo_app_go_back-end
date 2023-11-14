package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hemiknk/todo_app_go_back-end/internal/db"
	"github.com/hemiknk/todo_app_go_back-end/internal/handler"

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

	defer db.Conn.Close()

	http.HandleFunc("/edit", handler.EditHandler)
	http.HandleFunc("/delete", handler.DeteteItemHandler)
	http.HandleFunc("/create", handler.CreateHandler)
	http.HandleFunc("/", handler.RenderTemplate)

	log.Fatal(http.ListenAndServe("0.0.0.0:8018", nil))
}
