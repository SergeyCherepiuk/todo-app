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

func (repository TodoRepositoryImpl) Create(todo models.Todo) error {
	file, err := os.OpenFile("./src/database/db.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(todo.String()); err != nil {
		return err
	}
	return nil
}

func (repository TodoRepositoryImpl) Read() ([]models.Todo, error) {
	file, err := os.Open("./src/database/db.txt")
	if err != nil {
		return []models.Todo{}, nil
	}

	var fileContent []byte
	if fileStats, err := file.Stat(); err != nil {
		return []models.Todo{}, err
	} else {
		fileContent = make([]byte, fileStats.Size())
	}

	if _, err = file.Read(fileContent); err != nil {
		return []models.Todo{}, err
	}

	todos := []models.Todo{}
	for _, line := range strings.Split(string(fileContent), "\n") {
		todo, err := models.TodoFromString(line)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
