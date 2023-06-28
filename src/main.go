package main

import (
	"github.com/SergeyCherepiuk/todo-app/src/controllers"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	todoRepository := new(repositories.TodoRepositoryImpl)
	todoController := controllers.TodoContoller{Repository: *todoRepository}

	app.Get("/api/todo", todoController.ReadAll)
	app.Post("/api/todo", todoController.Create)

	app.Listen(":8000")
}
