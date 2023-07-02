package main

import (
	"github.com/SergeyCherepiuk/todo-app/src/controllers"
	"github.com/SergeyCherepiuk/todo-app/src/database"
	"github.com/SergeyCherepiuk/todo-app/src/initializers"
	"github.com/SergeyCherepiuk/todo-app/src/middleware"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	initializers.LoadEnv()
	db = database.Connect()
	database.Sync(db)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	authRepository := repositories.NewAuthRepository(db)
	authController := controllers.NewAuthController(authRepository)
	authMiddleware := middleware.NewAuthMiddleware()

	auth := api.Group("/auth")
	auth.Post("/signup", authController.SignUp)
	auth.Post("/login", authController.Login)
	auth.Post("/logout", authController.Logout)

	todoRepository := repositories.NewTodoRepository(db)
	todoController := controllers.NewTodoController(todoRepository)

	todo := api.Group("/todos")
	todo.Use(authMiddleware)
	todo.Get("/:todoId", todoController.GetById)
	todo.Get("/", todoController.GetAll)
	todo.Post("/", todoController.Create)
	todo.Put("/:todoId", todoController.Update)
	todo.Put("/toggle-completion/:todoId", todoController.ToggleCompletion)
	todo.Delete("/:todoId", todoController.Delete)
	todo.Delete("/", todoController.DeleteAll)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryController := controllers.NewCategoryController(categoryRepository)

	category := api.Group("/categories")
	category.Use(authMiddleware)
	category.Get("/:categoryId", categoryController.GetById)
	category.Get("/", categoryController.GetAll)
	category.Post("/", categoryController.Create)
	category.Put("/:categoryId", categoryController.Update)
	category.Delete("/:categoryId", categoryController.Delete)
	category.Delete("/", categoryController.DeleteAll)

	app.Listen(":8000")
}
