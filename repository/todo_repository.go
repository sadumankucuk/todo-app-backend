package repository

import "todo/model"

type ITodoRepository interface {
	CreateTodo(newTask model.TodoRequest) (*model.Todo, error)
	GetTodoList() (model.TodoResponse, error)
}

type TodoRepository struct {
	Todo map[int]*model.Todo
}

func NewITodoRepository() ITodoRepository {
	return &TodoRepository{
		Todo: map[int]*model.Todo{
			1: {
				ID:   1,
				Task: "go to the market",
			},
		},
	}
}

func (t TodoRepository) CreateTodo(newTask model.TodoRequest) (*model.Todo, error) {
	newTodoID := len(t.Todo) + 1
	t.Todo[newTodoID] = &model.Todo{
		ID:   newTodoID,
		Task: newTask.Task,
	}
	return t.Todo[newTodoID], nil
}

func (t TodoRepository) GetTodoList() (model.TodoResponse, error) {
	todoList := make([]model.Todo, 0, len(t.Todo))

	for _, task := range t.Todo {
		todoList = append(todoList, *task)
	}

	return todoList, nil
}
