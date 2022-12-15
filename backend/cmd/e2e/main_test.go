package main

import (
	"testing"

	"github.com/YunosukeY/kind-backend/internal/e2e"
	"github.com/stretchr/testify/assert"
)

var login = e2e.NewPostAPI[e2e.User, interface{}]("/api/v1/sessions")
var getTodos = e2e.NewGetAPI[[]e2e.TodoForResponse]("/api/v1/todos")
var postTodo = e2e.NewPostAPI[e2e.TodoForPostRequest, e2e.TodoForResponse]("/api/v1/todos")
var publicGetTodos = e2e.NewGetAPI[[]e2e.TodoForResponse]("/api/v1/public/todos")

func Test(t *testing.T) {
	testWithoutAuth(t)
	testWithAuth(t)
}

func testWithoutAuth(t *testing.T) {
	todos, err := publicGetTodos.Request()
	assert.Nil(t, err)
	assert.Equal(t, &[]e2e.TodoForResponse{}, todos)

	todos, err = getTodos.Request()
	assert.NotNil(t, err, todos)
}

func testWithAuth(t *testing.T) {
	_, err := login.Request(e2e.User{Name: "user", Password: "pass"})
	assert.Nil(t, err)

	expected := e2e.TodoForResponse{ID: 0, Content: "test"}

	todos, err := getTodos.Request()
	assert.Nil(t, err)
	assert.Equal(t, &[]e2e.TodoForResponse{}, todos)

	todo, err := postTodo.Request(e2e.TodoForPostRequest{Content: "test"})
	assert.Nil(t, err)
	assert.Equal(t, &expected, todo)

	todos, err = getTodos.Request()
	assert.Nil(t, err)
	assert.Equal(t, &[]e2e.TodoForResponse{expected}, todos)
}
