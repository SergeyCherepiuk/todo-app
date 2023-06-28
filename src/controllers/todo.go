package controllers

import (
	"net/http"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type TodoContoller struct {
	Repository repositories.TodoRepositoryImpl
}

func (controller TodoContoller) GetById(c *fiber.Ctx) error {
	// TODO: Implement
	return c.Status(http.StatusNotImplemented).JSON("message: Endpoint is not implemented yet")
}

func (controller TodoContoller) GetAll(c *fiber.Ctx) error {
	todos, err := controller.Repository.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON("message: Internal server error")
	}
	return c.Status(http.StatusOK).JSON(todos)
}

func (controller TodoContoller) Create(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusInternalServerError).JSON("message: Internal server error")
	}

	if todo.Title == "" || todo.Category == "" {
		return c.Status(http.StatusBadRequest).JSON("message: Not enough information provided")
	}

	if err := controller.Repository.Create(*todo); err != nil {
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
