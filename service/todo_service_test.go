package service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
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
		Return(&expectedTodo, nil).
		Times(1)

	service := service.NewITodoService(repository)
	newTodo, _ := service.CreateTodo(newTask)

	assert.Equal(t, &expectedTodo, newTodo)
}

func TestTodoService_GetTodoList(t *testing.T) {
	t.Run("should correctly return todoList", func(t *testing.T) {
		expectedTodoList := model.TodoResponse{
			{
				ID:   1,
				Task: "go to the market",
			},
			{
				ID:   2,
				Task: "buy some milk",
			},
		}

		repository := mock.NewMockITodoRepository(gomock.NewController(t))
		repository.EXPECT().
			GetTodoList().
			Return(expectedTodoList, nil).
			Times(1)

		service := service.NewITodoService(repository)
		todoList, _ := service.GetTodoList()

		assert.Equal(t, expectedTodoList, todoList)
	})
	t.Run("should correctly return sorted todoList", func(t *testing.T) {
		expectedTodoList := model.TodoResponse{
			{
				ID:   1,
				Task: "go to the market",
			},
			{
				ID:   4,
				Task: "test4",
			},
			{
				ID:   3,
				Task: "buy some milk",
			},
			{
				ID:   5,
				Task: "test5",
			},
			{
				ID:   2,
				Task: "test",
			},
		}

		deepCopy := model.TodoResponse{}
		copier.Copy(&deepCopy, &expectedTodoList)

		repository := mock.NewMockITodoRepository(gomock.NewController(t))
		repository.EXPECT().
			GetTodoList().
			Return(expectedTodoList, nil).
			Times(1)

		service := service.NewITodoService(repository)
		todoList, _ := service.GetTodoList()

		assert.NotEqual(t, deepCopy, todoList)
	})
	t.Run("should return error when repository error", func(t *testing.T) {

		expectedError := errors.New("repository error")
		repository := mock.NewMockITodoRepository(gomock.NewController(t))
		repository.EXPECT().
			GetTodoList().
			Return(nil, expectedError).
			Times(1)

		service := service.NewITodoService(repository)
		_, err := service.GetTodoList()

		assert.Equal(t, expectedError, err)
	})
}
