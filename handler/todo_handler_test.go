package handler_test

import (
	"bytes"
	"encoding/json"
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
	newTodo := model.Todo{ID: 1, Task: "buy some milk"}
	service := mock.NewMockITodoService(gomock.NewController(t))
	service.EXPECT().
		CreateTodo(model.TodoRequest{Task: "buy some milk"}).
		Return(&newTodo).
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

}
