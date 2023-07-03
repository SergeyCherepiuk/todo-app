package routes

import (
	"github.com/SergeyCherepiuk/todo-app/src/controllers"
	"github.com/SergeyCherepiuk/todo-app/src/middleware"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupTodoRoutes(api *fiber.Router, db *sqlx.DB) {
	todoRepository := repositories.NewTodoRepository(db)
	todoController := controllers.NewTodoController(todoRepository)

	authMiddleware := middleware.NewAuthMiddleware()
	ownerMiddleware := middleware.NewOwnerMiddleware("todoId", todoRepository.HasTodo)

	todo := (*api).Group("/todos")
	todo.Use(authMiddleware)
	todo.Get("/:todoId", ownerMiddleware, todoController.GetById)
	todo.Get("/", todoController.GetAll)
	todo.Post("/", todoController.Create)
	todo.Put("/:todoId", ownerMiddleware, todoController.Update)
	todo.Put("/toggle-completion/:todoId", ownerMiddleware, todoController.ToggleCompletion)
	todo.Delete("/:todoId", ownerMiddleware, todoController.Delete)
	todo.Delete("/", todoController.DeleteAll)
}
