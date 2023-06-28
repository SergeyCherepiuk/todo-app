package repositories

import (
	"os"
	"strings"

	"github.com/SergeyCherepiuk/todo-app/src/models"
)

type TodoRepository interface {
	Create(models.Todo) error
	Read() ([]models.Todo, error)
}

type TodoRepositoryImpl struct{}

// TODO: Add id parameter
func (repository TodoRepositoryImpl) GetById() (models.Todo, error) {
	//TODO: Implement
	return models.Todo{}, nil
}

func (repository TodoRepositoryImpl) GetAll() ([]models.Todo, error) {
	// TODO: Implement
	return []models.Todo{}, nil	
}

func (repository TodoRepositoryImpl) Create(todo models.Todo) error {
	// TODO: Implement
	return nil	
}

func (repository TodoRepositoryImpl) Update(todo models.Todo) error {
	// TODO: Implement
	return nil	
}

func (repository TodoRepositoryImpl) Delete(todo models.Todo) error {
	// TODO: Implement
	return nil
}