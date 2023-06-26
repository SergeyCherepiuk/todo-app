package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
)

type TodoContoller struct {
	Repository repositories.TodoRepositoryImpl
}

func (controller TodoContoller) Create(w http.ResponseWriter, req *http.Request) {
	todo := new(models.Todo)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	if err := json.NewDecoder(req.Body).Decode(todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("message: Internal server error")
		return
	}

	if todo.Title == "" || todo.Category.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode("message: Not enough information provided")
		return
	}

	if err := controller.Repository.Create(*todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode("message: Internal server error")
		return
	}
	
	w.WriteHeader(http.StatusOK)
	encoder.Encode("message: Todo successfully created")
}
