package controllers

import (
	"net/http"

	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	repository repositories.CategoryRepository
}

func NewCategoryController(repository repositories.CategoryRepository) *CategoryController {
	return &CategoryController{repository: repository}
}

func (controller CategoryController) GetById(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}

func (controller CategoryController) GetAll(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}

func (controller CategoryController) Create(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}

func (controller CategoryController) Update(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}

func (controller CategoryController) Delete(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}

func (controller CategoryController) DeleteAll(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}