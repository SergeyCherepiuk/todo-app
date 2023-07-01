package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
		return c.Status(http.StatusNotFound).JSON(messageResponse{
			Message: fmt.Sprintf("there is no todo with id = %d", id),
		})
	}
	return c.Status(http.StatusOK).JSON(todoResponse{Todo: todo})
}

func (controller TodoContoller) GetAll(c *fiber.Ctx) error {
	todos, err := controller.repository.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
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
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "invalid request body",
		})
	}

	if todo.Title == "" {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "not enough information provided",
		})
	}

	todo, err := controller.repository.Create(todo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: "invalid request body",
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

	decoder := json.NewDecoder(strings.NewReader(string(c.Body())))
	decoder.UseNumber()
	fieldsWithNewValues := make(map[string]any)
	err = decoder.Decode(&fieldsWithNewValues)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}

	updatedTodo, err := controller.repository.Update(id, fieldsWithNewValues)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(todoResponse{
		Todo: updatedTodo,
	})
}

func (controller TodoContoller) ToggleCompletion(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(utils.CopyString(c.Params("id")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not valid id", c.Params("id")),
		})
	}

	todo, err := controller.repository.ToggleCompletion(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(todoResponse{
		Todo: todo,
	})
}

func (controller TodoContoller) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(utils.CopyString(c.Params("id")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not valid id", c.Params("id")),
		})
	}

	err = controller.repository.Delete(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(messageResponse{
		Message: "todo deleted successfully",
	})
}

func (controller TodoContoller) DeleteAll(c *fiber.Ctx) error {
	nRows, err := controller.repository.DeleteAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(messageResponse{
		Message: fmt.Sprintf("%d todos deleted successfully", nRows),
	})
}
