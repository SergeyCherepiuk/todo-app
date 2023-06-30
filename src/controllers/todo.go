package controllers

import (
	"encoding/json"
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

type messageResponse struct {
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
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not valid id", c.Params("id")),
		})
	}

	todo, err := controller.repository.GetById(id)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusNotFound).JSON(messageResponse{
			Message: fmt.Sprintf("There is no todo with id = %d", id),
		})
	}
	return c.Status(http.StatusOK).JSON(todoResponse{Todo: todo})
}

func (controller TodoContoller) GetAll(c *fiber.Ctx) error {
	todos, err := controller.repository.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: "Internal server error",
		})
	}

	if len(todos) < 1 {
		c.Status(http.StatusNoContent)
	} else {
		c.Status(http.StatusOK)
	}
	return c.JSON(todosResponse{Todos: todos})
}

func (controller TodoContoller) Create(c *fiber.Ctx) error {
	todo := models.Todo{}
	if err := c.BodyParser(&todo); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "Invalid request body",
		})
	}

	if todo.Title == "" || todo.Category == "" {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "Not enough information provided",
		})
	}

	todo, err := controller.repository.Create(todo)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: "Invalid request body",
		})
	}

	return c.Status(http.StatusCreated).JSON(todoResponse{Todo: todo})
}

func (controller TodoContoller) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(utils.CopyString(c.Params("id")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not valid id", c.Params("id")),
		})
	}

	fieldsWithNewValues := make(map[string]any)
	err = json.Unmarshal(c.Body(), &fieldsWithNewValues)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: "Internal server error",
		})
	}

	updatedTodo, err := controller.repository.Update(id, fieldsWithNewValues)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: "Internal server error",
		})
	}
	return c.Status(http.StatusOK).JSON(todoResponse{
		Todo: updatedTodo,
	})
}

func (controller TodoContoller) Delete(c *fiber.Ctx) error {
	// TODO: Implement
	return c.Status(http.StatusNotImplemented).JSON(messageResponse{
		Message: "Endpoint is not implemented yet",
	})
}

func (controller TodoContoller) DeleteAll(c *fiber.Ctx) error {
	// TODO: Implement
	return c.Status(http.StatusNotImplemented).JSON(messageResponse{
		Message: "Endpoint is not implemented yet",
	})
}
