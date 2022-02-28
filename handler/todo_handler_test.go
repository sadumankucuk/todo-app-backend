//go:build unit_test
// +build unit_test

package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/handler"
	"todo/mock"
	"todo/model"
)

func TestTodoHandler_CreateTodo(t *testing.T) {
	t.Run("should return successfully newTodo when add todo", func(t *testing.T) {
		newTodo := model.Todo{ID: 1, Task: "buy some milk"}
		service := mock.NewMockITodoService(gomock.NewController(t))
		service.EXPECT().
			CreateTodo(model.TodoRequest{Task: "buy some milk"}).
			Return(&newTodo, nil).
			Times(1)

		handler := handler.NewITodoHandler(service)
		reqBody := model.TodoRequest{
			Task: "buy some milk",
		}

		payload, _ := json.Marshal(reqBody)
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(payload))
		w := httptest.NewRecorder()

		handler.CreateTodo(w, r)
		resBody, _ := json.Marshal(&newTodo)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", w.Result().Header.Get("content-type"))
		assert.Equal(t, string(resBody), w.Body.String())
	})
	t.Run("should return error when service error", func(t *testing.T) {
		service := mock.NewMockITodoService(gomock.NewController(t))
		service.EXPECT().
			CreateTodo(model.TodoRequest{Task: "buy some milk"}).
			Return(nil, errors.New("service error")).
			Times(1)

		handler := handler.NewITodoHandler(service)
		reqBody := model.TodoRequest{
			Task: "buy some milk",
		}

		payload, _ := json.Marshal(reqBody)
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(payload))
		w := httptest.NewRecorder()

		handler.CreateTodo(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestTodoHandler_GetTodoList(t *testing.T) {
	t.Run("should return correctly todoList", func(t *testing.T) {
		service := mock.NewMockITodoService(gomock.NewController(t))
		service.EXPECT().
			GetTodoList().
			Return(model.TodoResponse{
				{
					ID:   1,
					Task: "buy some milk",
				},
			}, nil).
			Times(1)

		handler := handler.NewITodoHandler(service)
		r := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		w := httptest.NewRecorder()

		handler.GetTodoList(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", w.Result().Header.Get("content-type"))

		expectedResBody := model.TodoResponse{}
		err := json.Unmarshal(w.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, errors.New("Error"))
		assert.Equal(t, expectedResBody[0].ID, 1)
	})
	t.Run("should return error when service error", func(t *testing.T) {
		service := mock.NewMockITodoService(gomock.NewController(t))
		service.EXPECT().
			GetTodoList().
			Return(nil, errors.New("service error")).
			Times(1)

		handler := handler.NewITodoHandler(service)
		r := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		w := httptest.NewRecorder()

		handler.GetTodoList(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
