package service

import (
	"sort"
	"todo/model"
	"todo/repository"
)

type ITodoService interface {
	CreateTodo(newTask model.TodoRequest) (*model.Todo, error)
	GetTodoList() (model.TodoResponse, error)
}

type TodoService struct {
	Repository repository.ITodoRepository
}

func NewITodoService(r repository.ITodoRepository) ITodoService {
	return &TodoService{Repository: r}
}

func (t TodoService) CreateTodo(newTask model.TodoRequest) (*model.Todo, error) {
	return t.Repository.CreateTodo(newTask)
}

func (t TodoService) GetTodoList() (model.TodoResponse, error) {
	todoList, _ := t.Repository.GetTodoList()
	sort.Slice(todoList, func(i, j int) bool {
		return todoList[i].ID < todoList[j].ID
	})
	return todoList, nil
}
