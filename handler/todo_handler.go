package handler

import (
	"net/http"
	"todo/service"
)

type ITodoHandler interface {
	CreateTodo(w http.ResponseWriter, r *http.Request)
}

type TodoHandler struct {
	Service service.ITodoService
}

func NewITodoHandler(service service.ITodoService) ITodoHandler {
	return &TodoHandler{Service: service}
}

func (t TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
