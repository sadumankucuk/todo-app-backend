package service

import (
	"todo/model"
	"todo/repository"
)

type ITodoService interface {
	CreateTodo(newTask model.TodoRequest) *model.Todo
	GetTodoList() model.TodoResponse
}

type TodoService struct {
	Repository repository.ITodoRepository
}

func NewITodoService(r repository.ITodoRepository) ITodoService {
	return &TodoService{Repository: r}
}

func (t TodoService) CreateTodo(newTask model.TodoRequest) *model.Todo {
	return t.Repository.CreateTodo(newTask)
}

func (t TodoService) GetTodoList() model.TodoResponse {
	//TODO implement me
	panic("implement me")
}
