package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	repository repositories.AuthRepository
}

func NewAuthController(repository repositories.AuthRepository) *AuthController {
	return &AuthController{repository: repository}
}

func (controller AuthController) SignUp(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "invalid request body",
		})
	}

	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "not enought information provided",
		})
	}

	signedToken, err := controller.repository.SignUp(user)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(messageResponse{
			Message: err.Error(),
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = signedToken
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)
	c.Cookie(cookie)
	return c.Status(http.StatusOK).JSON(messageResponse{
		Message: "signed up successfully",
	})
}

func (contoller AuthController) Login(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "invalid request body",
		})
	}
	
	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "not enought information provided",
		})
	}

	signedToken, err := contoller.repository.Login(user)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(messageResponse{
			Message: err.Error(),
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = signedToken
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)
	c.Cookie(cookie)
	return c.Status(http.StatusOK).JSON(messageResponse{
		Message: "logged in successfully",
	})
}

func (contoller AuthController) Logout(c *fiber.Ctx) error {
	if strings.TrimSpace(c.Cookies("token", "")) == "" {
		return c.Status(http.StatusBadRequest).JSON(messageResponse{
			Message: "you are not logged in",
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	c.Cookie(cookie)
	return c.Status(http.StatusOK).JSON(messageResponse{
		Message: "logged out successfully",
	})
}