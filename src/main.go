package main

import (
	"net/http"

	"github.com/SergeyCherepiuk/todo-app/src/controllers"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
)

func main() {
	todoRepository := new(repositories.TodoRepositoryImpl)
	todoController := controllers.TodoContoller{Repository: *todoRepository}

	http.HandleFunc("/api/todo/create", todoController.Create)
	http.HandleFunc("/api/todo/read", todoController.ReadAll)

	http.ListenAndServe("localhost:8000", nil) // handler=nil -> DefaultServeMux is used
}
