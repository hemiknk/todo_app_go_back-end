package model

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hemiknk/todo_app_go_back-end/internal/db"
)

func TestSaveTodoItem_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	db.Conn = mockDB

	expectedQuery := `INSERT INTO todo (title, done) VALUES (?, ?);`
	mock.ExpectExec(regexp.QuoteMeta(expectedQuery)).
		WithArgs("Test Item", true).
		WillReturnResult(sqlmock.NewResult(1, 1))

	item := ToDoItem{
		Title: "Test Item",
		Done:  true,
	}

	err = SaveTodoItem(item)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestSaveTodoItem_ErrorHandling(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	db.Conn = mockDB

	expectedQuery := `INSERT INTO todo (title, done) VALUES (?, ?);`
	mock.ExpectExec(regexp.QuoteMeta(expectedQuery)).
		WithArgs("Test Item", true).
		WillReturnError(fmt.Errorf("error saving todo item"))

	item := ToDoItem{
		Title: "Test Item",
		Done:  true,
	}

	err = SaveTodoItem(item)

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
