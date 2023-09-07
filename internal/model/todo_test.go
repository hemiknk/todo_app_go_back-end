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

func TestTodoList_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	db.Conn = mockDB

	expectedQuery := `SELECT id, title, done FROM todo;`
	mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "done"}).
			AddRow(1, "Test Item", true).
			AddRow(2, "Test Item 2", false))

	list, err := TodoList()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(list))
	assert.Equal(t, "Test Item", list[0].Title)
	assert.Equal(t, true, list[0].Done)
	assert.Equal(t, "Test Item 2", list[1].Title)
	assert.Equal(t, false, list[1].Done)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestTodoList_ErrorHandling(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	db.Conn = mockDB

	expectedQuery := `SELECT id, title, done FROM todo;`
	mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WillReturnError(fmt.Errorf("error getting todo list"))

	_, err = TodoList()

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
