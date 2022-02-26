package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"todo/handler"
	"todo/repository"
	"todo/service"
)

func main() {
	repository := repository.NewITodoRepository()
	service := service.NewITodoService(repository)
	handler := handler.NewITodoHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/todos", handler.CreateTodo).Methods(http.MethodPost)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println(err)
	}
}
