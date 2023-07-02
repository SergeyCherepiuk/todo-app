package repositories

import (
	"fmt"
	"strings"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/jmoiron/sqlx"
)

type TodoRepository interface {
	GetById(uint64) (models.Todo, error)
	GetAll() ([]models.Todo, error)
	Create(models.Todo) (models.Todo, error)
	Update(uint64, map[string]any) (models.Todo, error)
	ToggleCompletion(uint64) (models.Todo, error)
	Delete(uint64) error
	DeleteAll() (uint64, error)
}

type TodoRepositoryImpl struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{db: db}
}

func (repository TodoRepositoryImpl) GetById(id uint64) (models.Todo, error) {
	todo := models.Todo{}
	sql := "SELECT * FROM todos WHERE id = $1"
	err := repository.db.Get(&todo, sql, id)
	return todo, err
}

func (repository TodoRepositoryImpl) GetAll() ([]models.Todo, error) {
	todos := []models.Todo{}
	sql := "SELECT * FROM todos"
	err := repository.db.Select(&todos, sql)
	return todos, err
}

func (repository TodoRepositoryImpl) Create(todo models.Todo) (models.Todo, error) {
	insertedTodo := models.Todo{}
	sql := "INSERT INTO todos (title, category_id, priority, is_completed) VALUES ($1, $2, $3, $4) RETURNING *"
	row := repository.db.QueryRowx(sql, todo.Title, todo.Category, todo.Priority, todo.IsCompleted)
	if row.Err() != nil { 
		return insertedTodo, row.Err()
	}
	err := row.StructScan(&insertedTodo)
	return insertedTodo, err
}

func (repository TodoRepositoryImpl) Update(id uint64, fieldsWithNewValues map[string]any) (models.Todo, error) {
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

	row := repository.db.QueryRowx(string(sql), id)
	if row.Err() != nil {
		return updatedTodo, row.Err()
	}
	err := row.StructScan(&updatedTodo)
	return updatedTodo, err
}

func (repository TodoRepositoryImpl) ToggleCompletion(id uint64) (models.Todo, error) {
	todo := models.Todo{}
	sql := "UPDATE todos SET is_completed = NOT is_completed WHERE id = $1 RETURNING *"
	row := repository.db.QueryRowx(sql, id)
	if row.Err() != nil {
		return todo, row.Err()
	}
	err := row.StructScan(&todo)
	return todo, err
}

func (repository TodoRepositoryImpl) Delete(id uint64) error {
	sql := "DELETE FROM todos WHERE id = $1"
	res, err := repository.db.Exec(sql, id)
	if err != nil {
		return err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if nRows < 1 {
		return fmt.Errorf("todo with id = %d not found", id)
	}
	return nil
}

func (repository TodoRepositoryImpl) DeleteAll() (uint64, error) {
	sql := "DELETE FROM todos"
	res, err := repository.db.Exec(sql)
	if err != nil { 
		return 0, err
	}
	nRows, err := res.RowsAffected()
	if err != nil { 
		return 0, err
	}
	return uint64(nRows), nil
}
