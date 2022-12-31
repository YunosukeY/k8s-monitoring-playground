package main

import (
	"testing"

	"github.com/YunosukeY/kind-backend/internal/e2e"
	"github.com/stretchr/testify/assert"
)

var publicGetTodos = e2e.NewGetAPI[[]e2e.TodoForResponse]("/api/v1/public/todos")

func Test(t *testing.T) {
	testWithoutAuth(t)
}

func testWithoutAuth(t *testing.T) {
	todos, err := publicGetTodos.Request()
	assert.Nil(t, err)
	assert.Equal(t, &[]e2e.TodoForResponse{}, todos)
}
