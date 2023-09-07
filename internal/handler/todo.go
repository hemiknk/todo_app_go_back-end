package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/hemiknk/todo_app_go_back-end/internal/model"
)

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

	err = t.Execute(w, nil)
	if err != nil {
		log.Println("execute template error", err)
	}
}
