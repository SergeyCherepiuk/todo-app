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
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: fmt.Sprintf("Invalid id provided: %s", c.Params("categoryId")),
		})
	}

	category, err := controller.repository.GetById(categoryId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(MessageResponse{
			Message: fmt.Sprintf("Internal server error: %s", err.Error()),
		})
	}
	
	return c.Status(http.StatusOK).JSON(CategoryResponse{
		Category: category,
	})
}

func (controller CategoryController) GetAll(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint64)
	categories, err := controller.repository.GetAll(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(MessageResponse{
			Message: fmt.Sprintf("Internal server error: %s", err.Error()),
		})
	}

	if len(categories) < 1 {
		c.Status(http.StatusNoContent)
	} else {
		c.Status(http.StatusOK)
	}
	return c.JSON(CategoriesResponse{
		Categories: categories,
	})
}

func (controller CategoryController) Create(c *fiber.Ctx) error {
	category := models.Category{}
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: "invalid request body",
		})
	}

	userId := c.Locals("userId").(uint64)
	category.UserID = userId

	if strings.TrimSpace(category.Name) == "" {
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: "Not enough information provided (category name is missing)",
		})
	}

	insertedCategory, err := controller.repository.Create(category)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(MessageResponse{
			Message: fmt.Sprintf("Internal server error: %s", err.Error()),
		})
	}

	return c.Status(http.StatusCreated).JSON(CategoryResponse{
		Category: insertedCategory,
	})
}

func (controller CategoryController) Update(c *fiber.Ctx) error {
	categoryId, err := strconv.ParseUint(utils.CopyString(c.Params("categoryId")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: fmt.Sprintf("Invalid id provided: %s", c.Params("categoryId")),
		})
	}

	fieldsWithNewValues := make(map[string]any)
	decoder := json.NewDecoder(strings.NewReader(string(c.Body())))
	decoder.UseNumber()
	if err := decoder.Decode(&fieldsWithNewValues); err != nil {
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
	}

	updatedCategory, err := controller.repository.Update(categoryId, fieldsWithNewValues)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(MessageResponse{
			Message: fmt.Sprintf("Internal server error: %s", err.Error()),
		})
	}

	return c.Status(http.StatusOK).JSON(CategoryResponse{
		Category: updatedCategory,
	})
}

func (controller CategoryController) Delete(c *fiber.Ctx) error {
	categoryId, err := strconv.ParseUint(utils.CopyString(c.Params("categoryId")), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: fmt.Sprintf("Invalid id provided: %s", c.Params("categoryId")),
		})
	}

	if err := controller.repository.Delete(categoryId); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(MessageResponse{
			Message: fmt.Sprintf("Internal server error: %s", err.Error()),
		})
	}

	return c.Status(http.StatusOK).JSON(MessageResponse{
		Message: "Category deleted successfully",
	})
}

func (controller CategoryController) DeleteAll(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint64)
	nRows, err := controller.repository.DeleteAll(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(MessageResponse{
			Message: fmt.Sprintf("Internal server error: %s", err.Error()),
		})
	}

	return c.Status(http.StatusOK).JSON(MessageResponse{
		Message: fmt.Sprintf("%d category(-ies) deleted successfully", nRows),
	})
}
