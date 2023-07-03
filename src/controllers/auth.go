package controllers

import (
	"fmt"
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

func createCookie(c *fiber.Ctx, signedToken string, expiresIn time.Duration) {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = signedToken
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(expiresIn)
	c.Cookie(cookie)
}

func removeCookie(c *fiber.Ctx) {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Expires = time.Now()
	c.Cookie(cookie)
}

func (controller AuthController) SignUp(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
	}

	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: "Not enought information provided (username or password is missing)",
		})
	}

	signedToken, err := controller.repository.SignUp(user)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(MessageResponse{
			Message: fmt.Sprintf("Unauthorized: %s", err.Error()),
		})
	}

	createCookie(c, signedToken, 7*24*time.Hour)
	return c.Status(http.StatusOK).JSON(MessageResponse{
		Message: "Signed up successfully",
	})
}

func (contoller AuthController) Login(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
	}

	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return c.Status(http.StatusBadRequest).JSON(MessageResponse{
			Message: "Not enought information provided (username or password is missing)",
		})
	}

	signedToken, err := contoller.repository.Login(user)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(MessageResponse{
			Message: fmt.Sprintf("Unauthorized: %s", err.Error()),
		})
	}

	createCookie(c, signedToken, 7*24*time.Hour)
	return c.Status(http.StatusOK).JSON(MessageResponse{
		Message: "Logged in successfully",
	})
}

func (contoller AuthController) Logout(c *fiber.Ctx) error {
	if strings.TrimSpace(c.Cookies("token", "")) == "" {
		return c.Status(http.StatusUnauthorized).JSON(MessageResponse{
			Message: "You are not logged in",
		})
	}

	removeCookie(c)
	return c.Status(http.StatusOK).JSON(MessageResponse{
		Message: "Logged out successfully",
	})
}
