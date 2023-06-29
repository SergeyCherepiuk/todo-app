package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type TodoContoller struct {
	repository repositories.TodoRepository
}

func NewTodoController(repository repositories.TodoRepository) *TodoContoller {
	return &TodoContoller{repository: repository}
}

type errorMessageResponse struct {
	Message string `json:"message"`
}

type todoResponse struct {
	Todo models.Todo `json:"todo"`
}

type todosResponse struct {
	Todos []models.Todo `json:"todos"`
}

func (controller TodoContoller) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(utils.CopyString(c.Params("id")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(errorMessageResponse{
			Message: fmt.Sprintf("%s is not valid id", c.Params("id")),
		})
	}

	todo, err := controller.repository.GetById(id)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusNotFound).JSON(errorMessageResponse{
			Message: fmt.Sprintf("There is no todo with id = %d", id),
		})
	}
	return c.Status(http.StatusOK).JSON(todoResponse{Todo: todo})
}

func (controller TodoContoller) GetAll(c *fiber.Ctx) error {
	todos, err := controller.repository.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(errorMessageResponse{
			Message: "Internal server error",
		})
	}
	return c.Status(http.StatusOK).JSON(todosResponse{Todos: todos})
}

func (controller TodoContoller) Create(c *fiber.Ctx) error {
	todo := models.Todo{}
	if err := c.BodyParser(&todo); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(errorMessageResponse{
			Message: "Invalid request body",
		})
	}

	if todo.Title == "" || todo.Category == "" {
		return c.Status(http.StatusBadRequest).JSON(errorMessageResponse{
			Message: "Not enough information provided",
		})
	}

	todo, err := controller.repository.Create(todo)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(errorMessageResponse{
			Message: "Invalid request body",
		})
	}

	return c.Status(http.StatusCreated).JSON(todoResponse{Todo: todo})
}

func (controller TodoContoller) Update(c *fiber.Ctx) error {
	// TODO: Implement
	return c.Status(http.StatusNotImplemented).JSON(errorMessageResponse{
		Message: "Endpoint is not implemented yet",
	})
}

func (controller TodoContoller) Delete(c *fiber.Ctx) error {
	// TODO: Implement
	return c.Status(http.StatusNotImplemented).JSON(errorMessageResponse{
		Message: "Endpoint is not implemented yet",
	})
}
