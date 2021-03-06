// Code generated by MockGen. DO NOT EDIT.
// Source: service/todo_service.go

// Package mock_service is a generated GoMock package.
package mock

import (
	reflect "reflect"
	model "todo/model"

	gomock "github.com/golang/mock/gomock"
)

// MockITodoService is a mock of ITodoService interface.
type MockITodoService struct {
	ctrl     *gomock.Controller
	recorder *MockITodoServiceMockRecorder
}

// MockITodoServiceMockRecorder is the mock recorder for MockITodoService.
type MockITodoServiceMockRecorder struct {
	mock *MockITodoService
}

// NewMockITodoService creates a new mock instance.
func NewMockITodoService(ctrl *gomock.Controller) *MockITodoService {
	mock := &MockITodoService{ctrl: ctrl}
	mock.recorder = &MockITodoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITodoService) EXPECT() *MockITodoServiceMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockITodoService) CreateTodo(newTask model.TodoRequest) (*model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", newTask)
	ret0, _ := ret[0].(*model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockITodoServiceMockRecorder) CreateTodo(newTask interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockITodoService)(nil).CreateTodo), newTask)
}

// GetTodoList mocks base method.
func (m *MockITodoService) GetTodoList() (model.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoList")
	ret0, _ := ret[0].(model.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoList indicates an expected call of GetTodoList.
func (mr *MockITodoServiceMockRecorder) GetTodoList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoList", reflect.TypeOf((*MockITodoService)(nil).GetTodoList))
}
