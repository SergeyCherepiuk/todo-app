package repositories

import (
	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/jmoiron/sqlx"
)

type TodoRepository interface {
	GetById(uint64) (models.Todo, error)
	GetAll() ([]models.Todo, error)
	Create(models.Todo) (models.Todo, error)
	Update(uint64) (models.Todo, error)
	Delete(uint64) error
}

type TodoRepositoryImpl struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{db: db}
}

// TODO: Add id parameter
func (repository TodoRepositoryImpl) GetById(id uint64) (models.Todo, error) {
	var todo models.Todo
	sql := "SELECT * FROM todo WHERE id = $1"
	row := repository.db.QueryRowx(sql, id)
	err := row.StructScan(&todo)
	return todo, err
}

func (repository TodoRepositoryImpl) GetAll() ([]models.Todo, error) {
	todos := []models.Todo{}
	sql := "SELECT * FROM todo"
	err := repository.db.Select(&todos, sql)
	return todos, err
}

func (repository TodoRepositoryImpl) Create(todo models.Todo) (models.Todo, error) {
	sql := "INSERT INTO todo (title, category, priority) VALUES ($1, $2, $3)"
	row := repository.db.QueryRowx(sql, todo.Title, todo.Category, todo.Priority)
	return todo, row.Err()
}

func (repository TodoRepositoryImpl) Update(uint64) (models.Todo, error) {
	// TODO: Implement
	return models.Todo{}, nil
}

func (repository TodoRepositoryImpl) Delete(uint64) error {
	// TODO: Implement
	return nil
}
