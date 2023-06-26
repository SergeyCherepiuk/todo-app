package repositories

import (
	"fmt"
	"os"

	"github.com/SergeyCherepiuk/todo-app/src/models"
)

type TodoRepository interface {
	Create(models.Todo)
}

type TodoRepositoryImpl struct{}

func (repository TodoRepositoryImpl) Create(todo models.Todo) error {
	file, err := os.OpenFile("./src/database/db.txt", os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	str := fmt.Sprintf("%s, %s, %d\n", todo.Title, todo.Category.Name, todo.Priority)
	if _, err := file.WriteString(str); err != nil {
		return err
	}
	return nil
}
