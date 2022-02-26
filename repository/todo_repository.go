package repository

import "todo/model"

type ITodoRepository interface {
	CreateTodo(newTask model.TodoRequest) *model.Todo
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

func (t TodoRepository) CreateTodo(newTask model.TodoRequest) *model.Todo {
	newTodoID := len(t.Todo) + 1
	t.Todo[newTodoID] = &model.Todo{
		ID:   newTodoID,
		Task: newTask.Task,
	}
	return t.Todo[newTodoID]
}