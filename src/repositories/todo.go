package repositories

import (
	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/jmoiron/sqlx"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	Create(models.Todo) (uint64, error)
}

type TodoRepositoryImpl struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{db: db}
}

// TODO: Add id parameter
func (repository TodoRepositoryImpl) GetById() (models.Todo, error) {
	//TODO: Implement
	return models.Todo{}, nil
}

func (repository TodoRepositoryImpl) GetAll() ([]models.Todo, error) {
	todos := []models.Todo{}
	sql := "SELECT * FROM todo"
	if err := repository.db.Select(&todos, sql); err != nil {
		return todos, err
	}
	return todos, nil
}

func (repository TodoRepositoryImpl) Create(todo models.Todo) (uint64, error) {
	var id uint64
	sql := "INSERT INTO todo (title, category, priority) VALUES ($1, $2, $3) RETURNING id"
	row := repository.db.QueryRow(sql, todo.Title, todo.Category, todo.Priority)
	err := row.Scan(&id)
	return id, err
}

func (repository TodoRepositoryImpl) Update(todo models.Todo) error {
	// TODO: Implement
	return nil
}

func (repository TodoRepositoryImpl) Delete(todo models.Todo) error {
	// TODO: Implement
	return nil
}
