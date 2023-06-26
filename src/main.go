package main

import (
	"fmt"
	"net/http"

	"github.com/SergeyCherepiuk/todo-app/src/controllers"
	"github.com/SergeyCherepiuk/todo-app/src/repositories"
)

func pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "pong")
}

func main() {
	todoRepository := new(repositories.TodoRepositoryImpl)
	todoController := controllers.TodoContoller{Repository: *todoRepository}

	http.HandleFunc("/ping", pong)	
	http.HandleFunc("/api/todo/create", todoController.Create)

	http.ListenAndServe("localhost:8000", nil) // handler=nil -> DefaultServeMux is used
}
