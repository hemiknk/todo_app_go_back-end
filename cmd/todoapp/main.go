package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/hemiknk/todo_app_go_back-end/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

type ToDoItem struct {
	ID    int
	Title string
	Done  bool
}

var conn *sql.DB

func saveTodoItem(item ToDoItem) error {
	query := `
		INSERT INTO todo (title, done) VALUES (?, ?);
	`
	_, err := conn.Exec(query, item.Title, item.Done)
	if err != nil {
		return err
	}

	return nil
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	item := ToDoItem{Title: title, Done: false}

	err := saveTodoItem(item)
	if err != nil {
		log.Println("can't save todo item", err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/template/index.html")
	if err != nil {
		log.Println("parse template error", err)
	}

	w.Header().Set("Content-Type", "text/html")

	err = t.Execute(w, nil)
	if err != nil {
		log.Println("execute template error", err)
	}
}

func main() {
	dbConn, err := db.GetConnection()
	if err != nil {
		panic(err)
	}
	conn = dbConn

	defer conn.Close()

	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/", renderTemplate)

	log.Fatal(http.ListenAndServe("localhost:8018", nil))
}
