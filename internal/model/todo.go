package model

import (
	"fmt"
	"github.com/hemiknk/todo_app_go_back-end/internal/db"
)

type ToDoItem struct {
	ID    int
	Title string
	Done  bool
}

// use sqlmock.Sqlmock to mock requests to database check that correct query is sent

func SaveTodoItem(item ToDoItem) error {
	query := `
		INSERT INTO todo (title, done) VALUES (?, ?);
	`

	_, err := db.Conn.Exec(query, item.Title, item.Done)

	if err != nil {
		return fmt.Errorf("error saving todo item: %v", err)
	}

	return nil
}
