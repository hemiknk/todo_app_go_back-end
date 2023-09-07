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

func MarkItemAsDone(id string) error {
	query := `
		UPDATE todo SET done = true WHERE id = ?;
	`

	_, err := db.Conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error marking item as done: %v", err)
	}

	return nil
}

func DeleteItem(id string) error {
	query := `
		DELETE FROM todo WHERE id = ?;
	`

	_, err := db.Conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting todo item: %v", err)
	}

	return nil
}

func TodoList() ([]ToDoItem, error) {
	query := `
		SELECT id, title, done FROM todo;
	`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting todo list: %v", err)
	}

	defer rows.Close()

	var list []ToDoItem

	for rows.Next() {
		var item ToDoItem
		err := rows.Scan(&item.ID, &item.Title, &item.Done)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		list = append(list, item)
	}

	return list, nil
}

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
