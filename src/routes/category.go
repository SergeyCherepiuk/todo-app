package routes

import (
	"github.com/SergeyCherepiuk/todo-app/src/controllers"
	"github.com/SergeyCherepiuk/todo-app/src/middleware"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupCategoryRoutes(api *fiber.Router, db *sqlx.DB) {
	categoryRepository := repositories.NewCategoryRepository(db)
	categoryController := controllers.NewCategoryController(categoryRepository)

	authMiddleware := middleware.NewAuthMiddleware()
	ownerMiddleware := middleware.NewOwnerMiddleware("categoryId", categoryRepository.HasCategory)

	category := (*api).Group("/categories")
	category.Use(authMiddleware)
	category.Get("/:categoryId", ownerMiddleware, categoryController.GetById)
	category.Get("/", categoryController.GetAll)
	category.Post("/", categoryController.Create)
	category.Put("/:categoryId", ownerMiddleware, categoryController.Update)
	category.Delete("/:categoryId", ownerMiddleware, categoryController.Delete)
	category.Delete("/", categoryController.DeleteAll)
}
