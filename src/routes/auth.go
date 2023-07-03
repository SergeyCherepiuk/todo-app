package routes

import (
	"github.com/SergeyCherepiuk/todo-app/src/controllers"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupAuthRoutes(api *fiber.Router, db *sqlx.DB) {
	authRepository := repositories.NewAuthRepository(db)
	authController := controllers.NewAuthController(authRepository)

	auth := (*api).Group("/auth")
	auth.Post("/signup", authController.SignUp)
	auth.Post("/login", authController.Login)
	auth.Post("/logout", authController.Logout)
}
