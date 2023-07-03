package repositories

import (
	"fmt"
	"strings"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/jmoiron/sqlx"
)

type TodoRepository interface {
	HasTodo(userId, todoId uint64) (bool, error)
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

func (repository TodoRepositoryImpl) HasTodo(todoId, userId uint64) (bool, error) {
	query := `SELECT COUNT(*) FROM todos WHERE id = :todoId AND user_id = :userId`
	namedParams := map[string]any{
		"todoId": todoId,
		"userId": userId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return false, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var count int
	if err := stmt.Get(&count, namedParams); err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}

	return count > 0, nil
}

func (repository TodoRepositoryImpl) GetById(todoId uint64) (models.Todo, error) {
	query := `SELECT * FROM todos WHERE id = :todoId`
	namedParams := map[string]any{
		"todoId": todoId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	todo := models.Todo{}
	if err := stmt.Get(&todo, namedParams); err != nil {
		return models.Todo{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return todo, nil
}

func (repository TodoRepositoryImpl) GetAll(userId uint64) ([]models.Todo, error) {
	query := `SELECT * FROM todos WHERE user_id = :userId`
	namedParams := map[string]any{
		"userId": userId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return []models.Todo{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	todos := []models.Todo{}
	if err := stmt.Select(&todos, namedParams); err != nil {
		return []models.Todo{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return todos, nil
}

func (repository TodoRepositoryImpl) Create(todo models.Todo) (models.Todo, error) {
	query := `INSERT INTO todos (title, priority, is_completed, user_id, category_id) VALUES (:title, :priority, :is_completed, :user_id, :category_id) RETURNING *`
	namedParams := map[string]any{
		"title":        todo.Title,
		"priority":     todo.Priority,
		"is_completed": todo.IsCompleted,
		"user_id":      todo.UserID,
		"category_id":  todo.CategoryID,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	insertedTodo := models.Todo{}
	if err := stmt.Get(&insertedTodo, namedParams); err != nil {
		return models.Todo{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return insertedTodo, nil
}

func (repository TodoRepositoryImpl) Update(todoId uint64, fieldsWithNewValues map[string]any) (models.Todo, error) {
	query := []byte("UPDATE todos SET ")
	namedParams := map[string]any{
		"todoId": todoId,
	}

	updates := []string{}
	for field, newValue := range fieldsWithNewValues {
		switch newValue.(type) {
		case string, byte, rune:
			updates = append(updates, fmt.Sprintf("%s = '%s'", field, newValue))
		default:
			updates = append(updates, fmt.Sprintf("%s = %s", field, newValue))
		}
	}
	query = append(query, strings.Join(updates, ", ")...)
	query = append(query, " WHERE id = :todoId RETURNING *"...)

	stmt, err := repository.db.PrepareNamed(string(query))
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	updatedTodo := models.Todo{}
	if err := stmt.Get(&updatedTodo, namedParams); err != nil {
		return models.Todo{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return updatedTodo, nil
}

func (repository TodoRepositoryImpl) ToggleCompletion(todoId uint64) (models.Todo, error) {
	query := `UPDATE todos SET is_completed = NOT is_completed WHERE id = :todoId RETURNING *`
	namedParams := map[string]any{
		"todoId": todoId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	todo := models.Todo{}
	if err := stmt.Get(&todo, namedParams); err != nil {
		return models.Todo{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return todo, nil
}

func (repository TodoRepositoryImpl) Delete(todoId uint64) error {
	query := `DELETE FROM todos WHERE id = :todoId`
	namedParams := map[string]any{
		"todoId": todoId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(namedParams)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get number of rows: %w", err)
	}
	if nRows < 1 {
		return fmt.Errorf("failed to find todo with id %d", todoId)
	}
	return nil
}

func (repository TodoRepositoryImpl) DeleteAll(userId uint64) (uint64, error) {
	query := `DELETE FROM todos WHERE user_id = :userId`
	namedParams := map[string]any{
		"userId": userId,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(namedParams)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get number of rows: %w", err)
	}
	return uint64(nRows), nil
}
