package controllers

import (
	"fmt"
	"net/http"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type TodoContoller struct {
	repository repositories.TodoRepository
}

func NewTodoController(repository repositories.TodoRepository) *TodoContoller {
	return &TodoContoller{repository: repository}
}

func (controller TodoContoller) GetById(c *fiber.Ctx) error {
	// TODO: Implement
	return c.Status(http.StatusNotImplemented).JSON("message: Endpoint is not implemented yet")
}

func (controller TodoContoller) GetAll(c *fiber.Ctx) error {
	todos, err := controller.repository.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON("message: Internal server error")
	}
	return c.Status(http.StatusOK).JSON(todos)
}

func (controller TodoContoller) Create(c *fiber.Ctx) error {
	todo := models.Todo{}
	if err := c.BodyParser(&todo); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON("message: Internal server error")
	}

	if todo.Title == "" || todo.Category == "" {
		return c.Status(http.StatusBadRequest).JSON("message: Not enough information provided")
	}

	if err := controller.repository.Create(todo); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON("message: Internal server error")
	}

	return c.Status(http.StatusOK).JSON("message: Todo successfully created")
}

func (controller TodoContoller) Update(c *fiber.Ctx) error {
	// TODO: Implement
	return c.Status(http.StatusNotImplemented).JSON("message: Endpoint is not implemented yet")
}

func (controller TodoContoller) Delete(c *fiber.Ctx) error {
	// TODO: Implement
	return c.Status(http.StatusNotImplemented).JSON("message: Endpoint is not implemented yet")
}
