package repository_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo/model"
	"todo/repository"
)

func TestTodoRepository_CreateTodo(t *testing.T) {
	expectedResponse := model.Todo{ID: 2, Task: "test"}
	repository := repository.NewITodoRepository()
	newTask := model.TodoRequest{
		Task: "test",
	}
	response := repository.CreateTodo(newTask)

	assert.Equal(t, &expectedResponse, response)
}

func TestTodoRepository_GetTodoList(t *testing.T) {
	repository := repository.NewITodoRepository()

	todoList := model.TodoResponse{
		{
			ID:   1,
			Task: "go to the market",
		},
		{
			ID:   2,
			Task: "buy some milk",
		},
	}
	repository.CreateTodo(model.TodoRequest{
		Task: "buy some milk",
	})
	response := repository.GetTodoList()
	lengthOfTodoList := len(repository.GetTodoList())
	assert.Equal(t, len(todoList), lengthOfTodoList)
	assert.Equal(t, todoList, response)
}
