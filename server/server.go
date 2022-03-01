package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"todo/handler"
	"todo/repository"
	"todo/service"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port string) error {
	repository := repository.NewITodoRepository()
	service := service.NewITodoService(repository)
	handler := handler.NewITodoHandler(service)

	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})
	routerHandler := c.Handler(r)

	r.HandleFunc("/api/v1/todos", handler.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/todos", handler.GetTodoList).Methods(http.MethodGet)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), routerHandler)
	return err
}
