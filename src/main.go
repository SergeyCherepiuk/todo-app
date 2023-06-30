package main

import (
	"github.com/SergeyCherepiuk/todo-app/src/controllers"
	"github.com/SergeyCherepiuk/todo-app/src/database"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	db = database.Connect()
	database.Sync(db)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	todoRepository := repositories.NewTodoRepository(db)
	todoController := controllers.NewTodoController(todoRepository)

	todo := api.Group("/todos")
	todo.Get("/:id", todoController.GetById)
	todo.Get("/", todoController.GetAll)
	todo.Post("/", todoController.Create)
	todo.Put("/:id", todoController.Update)
	todo.Delete("/:id", todoController.Delete)
	todo.Delete("/", todoController.DeleteAll)

	app.Listen(":8000")
}
