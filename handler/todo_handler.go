package handler

import (
	"encoding/json"
	"net/http"
	"todo/model"
	"todo/service"
)

type ITodoHandler interface {
	CreateTodo(w http.ResponseWriter, r *http.Request)
	GetTodoList(w http.ResponseWriter, r *http.Request)
}

type TodoHandler struct {
	Service service.ITodoService
}

func NewITodoHandler(service service.ITodoService) ITodoHandler {
	return &TodoHandler{Service: service}
}

func (t TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTask model.TodoRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&newTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo, _ := t.Service.CreateTodo(newTask)
	jsonBytes, err := json.Marshal(newTodo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	return
}

func (t TodoHandler) GetTodoList(w http.ResponseWriter, r *http.Request) {
	todoList := t.Service.GetTodoList()

	jsonBytes, err := json.Marshal(todoList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	return
}
