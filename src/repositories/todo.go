package repositories

import (
	"fmt"
	"strings"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/jmoiron/sqlx"
)

type TodoRepository interface {
	GetById(todoId uint64) (models.Todo, error)
	GetAll(userId uint64) ([]models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	Update(todoId uint64, fieldsWithNewValues map[string]any) (models.Todo, error)
	ToggleCompletion(todoId uint64) (models.Todo, error)
	Delete(todoId uint64) error
	DeleteAll(userId uint64) (uint64, error)
}

type TodoRepositoryImpl struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{db: db}
}

func (repository TodoRepositoryImpl) GetById(todoId uint64) (models.Todo, error) {
	todo := models.Todo{}
	sql := "SELECT * FROM todos WHERE id = $1"
	err := repository.db.Get(&todo, sql, todoId)
	return todo, err
}

func (repository TodoRepositoryImpl) GetAll(userId uint64) ([]models.Todo, error) {
	todos := []models.Todo{}
	sql := "SELECT * FROM todos WHERE user_id = $1"
	err := repository.db.Select(&todos, sql, userId)
	return todos, err
}

func (repository TodoRepositoryImpl) Create(todo models.Todo) (models.Todo, error) {
	insertedTodo := models.Todo{}
	sql := "INSERT INTO todos (title, priority, is_completed, user_id, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING *"
	row := repository.db.QueryRowx(sql, todo.Title, todo.Priority, todo.IsCompleted, todo.UserID, todo.Category)
	if row.Err() != nil { 
		return insertedTodo, row.Err()
	}
	err := row.StructScan(&insertedTodo)
	return insertedTodo, err
}

func (repository TodoRepositoryImpl) Update(todoId uint64, fieldsWithNewValues map[string]any) (models.Todo, error) {
	updatedTodo := models.Todo{}
	sql := []byte("UPDATE todos SET ")

	updates := []string{}
	for field, newValue := range fieldsWithNewValues {
		switch newValue.(type) {
		case string, byte, rune:
			updates = append(updates, fmt.Sprintf("%s = '%s'", field, newValue))
		default:
			updates = append(updates, fmt.Sprintf("%s = %s", field, newValue))
		}
	}
	sql = append(sql, strings.Join(updates, ", ")...)
	sql = append(sql, " WHERE id = $1 RETURNING *"...)

	row := repository.db.QueryRowx(string(sql), todoId)
	if row.Err() != nil {
		return updatedTodo, row.Err()
	}
	err := row.StructScan(&updatedTodo)
	return updatedTodo, err
}

func (repository TodoRepositoryImpl) ToggleCompletion(todoId uint64) (models.Todo, error) {
	todo := models.Todo{}
	sql := "UPDATE todos SET is_completed = NOT is_completed WHERE id = $1 RETURNING *"
	row := repository.db.QueryRowx(sql, todoId)
	if row.Err() != nil {
		return todo, row.Err()
	}
	err := row.StructScan(&todo)
	return todo, err
}

func (repository TodoRepositoryImpl) Delete(todoId uint64) error {
	sql := "DELETE FROM todos WHERE id = $1"
	res, err := repository.db.Exec(sql, todoId)
	if err != nil {
		return err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nRows < 1 {
		return fmt.Errorf("todo with id = %d not found", todoId)
	}
	return nil
}

func (repository TodoRepositoryImpl) DeleteAll(userId uint64) (uint64, error) {
	sql := "DELETE FROM todos WHERE user_id = $1"
	res, err := repository.db.Exec(sql, userId)
	if err != nil { 
		return 0, err
	}
	nRows, err := res.RowsAffected()
	if err != nil { 
		return 0, err
	}
	return uint64(nRows), nil
}
