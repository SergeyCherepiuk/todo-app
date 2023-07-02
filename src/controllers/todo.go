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

func (controller TodoContoller) GetById(c *fiber.Ctx) error {
	todoId, err := strconv.ParseUint(utils.CopyString(c.Params("todoId")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not a valid id", c.Params("todoId")),
		})
	}

	todo, err := controller.repository.GetById(todoId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(messageResponse{
			Message: fmt.Sprintf("there is no todo with id = %d", todoId),
		})
	}
	return c.Status(http.StatusOK).JSON(todoResponse{
		Todo: todo,
	})
}

func (controller TodoContoller) GetAll(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint64)

	todos, err := controller.repository.GetAll(userId)
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
	userId := c.Locals("userId").(uint64)

	todo := models.Todo{}
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "invalid request body",
		})
	}
	todo.UserID = userId

	if strings.TrimSpace(todo.Title) == "" {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "not enough information provided",
		})
	}

	insertedTodo, err := controller.repository.Create(todo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(todoResponse{
		Todo: insertedTodo,
	})
}

func (controller TodoContoller) Update(c *fiber.Ctx) error {
	todoId, err := strconv.ParseUint(utils.CopyString(c.Params("todoId")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not a valid id", c.Params("todoId")),
		})
	}

	fieldsWithNewValues := make(map[string]any)
	decoder := json.NewDecoder(strings.NewReader(string(c.Body())))
	decoder.UseNumber()
	if err := decoder.Decode(&fieldsWithNewValues); err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "invalid request body",
		})
	}

	updatedTodo, err := controller.repository.Update(todoId, fieldsWithNewValues)
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
	todoId, err := strconv.ParseUint(utils.CopyString(c.Params("todoId")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not a valid id", c.Params("todoId")),
		})
	}

	todo, err := controller.repository.ToggleCompletion(todoId)
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
	todoId, err := strconv.ParseUint(utils.CopyString(c.Params("todoId")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not a valid id", c.Params("todoId")),
		})
	}

	err = controller.repository.Delete(todoId)
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
	userId := c.Locals("userId").(uint64)

	nRows, err := controller.repository.DeleteAll(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(messageResponse{
		Message: fmt.Sprintf("%d todo(s) deleted successfully", nRows),
	})
}
