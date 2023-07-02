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

type CategoryController struct {
	repository repositories.CategoryRepository
}

func NewCategoryController(repository repositories.CategoryRepository) *CategoryController {
	return &CategoryController{repository: repository}
}

func (controller CategoryController) GetById(c *fiber.Ctx) error {
	categoryId, err := strconv.ParseUint(utils.CopyString(c.Params("categoryId")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not a valid id", c.Params("categoryId")),
		})
	}

	category, err := controller.repository.GetById(categoryId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(categoryResponse{
		Category: category,
	})
}

func (controller CategoryController) GetAll(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint64)

	categories, err := controller.repository.GetAll(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}

	if len(categories) < 1 {
		c.Status(http.StatusNoContent)
	} else {
		c.Status(http.StatusOK)
	}
	return c.JSON(categoriesResponse{
		Categories: categories,
	})
}

func (controller CategoryController) Create(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint64)

	category := models.Category{}
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "invalid request body",
		})
	}
	category.UserID = userId

	if strings.TrimSpace(category.Name) == "" {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "not enough information provided",
		})
	}

	insertedCategory, err := controller.repository.Create(category)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(categoryResponse{
		Category: insertedCategory,
	})
}

func (controller CategoryController) Update(c *fiber.Ctx) error {
	categoryId, err := strconv.ParseUint(utils.CopyString(c.Params("categoryId")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not a valid id", c.Params("categoryId")),
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

	updatedCategory, err := controller.repository.Update(categoryId, fieldsWithNewValues)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(categoryResponse{
		Category: updatedCategory,
	})
}

func (controller CategoryController) Delete(c *fiber.Ctx) error {
	categoryId, err := strconv.ParseUint(utils.CopyString(c.Params("categoryId")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: fmt.Sprintf("%s is not a valid id", c.Params("categoryId")),
		})
	}

	if err := controller.repository.Delete(categoryId); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(messageResponse{
		Message: "category deleted successfully",
	})
}

func (controller CategoryController) DeleteAll(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint64)

	nRows, err := controller.repository.DeleteAll(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(messageResponse{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(messageResponse{
		Message: fmt.Sprintf("%d category(-ies) deleted successfully", nRows),
	})
}
