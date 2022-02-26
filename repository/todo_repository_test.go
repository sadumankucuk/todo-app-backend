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
