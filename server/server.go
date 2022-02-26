package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"todo/handler"
	"todo/repository"
	"todo/service"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port int) error {
	repository := repository.NewITodoRepository()
	service := service.NewITodoService(repository)
	handler := handler.NewITodoHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/todos", handler.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/todos", handler.GetTodoList).Methods(http.MethodGet)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	return err
}
