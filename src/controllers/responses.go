package controllers

import "github.com/SergeyCherepiuk/todo-app/src/models"

type MessageResponse struct {
	Message string `json:"message"`
}

type TodoResponse struct {
	Todo models.Todo `json:"todo"`
}

type TodosResponse struct {
	Todos []models.Todo `json:"todos"`
}

type CategoryResponse struct {
	Category models.Category `json:"category"`
}

type CategoriesResponse struct {
	Categories []models.Category `json:"categories"`
}
