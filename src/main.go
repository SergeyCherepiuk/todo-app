package main

import (
	"github.com/SergeyCherepiuk/todo-app/src/database"
	"github.com/SergeyCherepiuk/todo-app/src/initializers"
	"github.com/SergeyCherepiuk/todo-app/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	initializers.LoadEnv()
	db = database.MustConnect()
	database.MustSync(db)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")
	routes.SetupAuthRoutes(&api, db)
	routes.SetupTodoRoutes(&api, db)
	routes.SetupCategoryRoutes(&api, db)

	app.Listen(":8000")
}
