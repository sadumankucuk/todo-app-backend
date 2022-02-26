package service

import (
	"sort"
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
	todoList := t.Repository.GetTodoList()
	sort.Slice(todoList, func(i, j int) bool {
		return todoList[i].ID < todoList[j].ID
	})
	return todoList
}
