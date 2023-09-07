package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/hemiknk/todo_app_go_back-end/internal/model"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
	// mark item as done
	id := r.URL.Query().Get("id")
	err := model.MarkItemAsDone(id)
	if err != nil {
		log.Println("can't mark item as done", err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	item := model.ToDoItem{Title: title, Done: false}

	err := model.SaveTodoItem(item)
	if err != nil {
		log.Println("can't save todo item", err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/template/index.html")
	if err != nil {
		log.Println("parse template error", err)
	}

	w.Header().Set("Content-Type", "text/html")

	items, err := model.TodoList()
	if err != nil {
		log.Println("can't get todo list", err)
	}
	err = t.Execute(w, items)
	if err != nil {
		log.Println("execute template error", err)
	}
}

func DeteteItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := model.DeleteItem(id)
	if err != nil {
		log.Println("can't delete todo item", err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
