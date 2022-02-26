package service_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo/mock"
	"todo/model"
	"todo/service"
)

func TestTodoService_CreateTodo(t *testing.T) {
	expectedTodo := model.Todo{ID: 1, Task: "buy some milk"}
	repository := mock.NewMockITodoRepository(gomock.NewController(t))
	newTask := model.TodoRequest{
		Task: "buy some milk",
	}
	repository.EXPECT().
		CreateTodo(newTask).
		Return(&expectedTodo).
		Times(1)

	service := service.NewITodoService(repository)
	newTodo := service.CreateTodo(newTask)

	assert.Equal(t, &expectedTodo, newTodo)
}
