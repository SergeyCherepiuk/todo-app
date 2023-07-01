package controllers

import "github.com/SergeyCherepiuk/todo-app/src/models"

type messageResponse struct {
	Message string `json:"message"`
}

type todoResponse struct {
	Todo models.Todo `json:"todo"`
}

type todosResponse struct {
	Todos []models.Todo `json:"todos"`
}

type categoryResponse struct {
	Category models.Category `json:"category"`
}

type categoriesResponse struct {
	Categories []models.Category `json:"categories"`
}
