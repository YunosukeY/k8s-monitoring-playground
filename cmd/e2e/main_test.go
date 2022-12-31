package main

import (
	"testing"
	"time"

	"github.com/YunosukeY/kind-backend/internal/e2e"
	"github.com/stretchr/testify/assert"
)

var publicGetTodos = e2e.NewGetAPI[[]e2e.TodoForResponse]("/api/v1/public/todos")
var getTraces = e2e.NewGetAPIWithUser[e2e.Traces]("admin", "admin", "localhost:3000", "/api/datasources/proxy/1/api/traces")

func Test(t *testing.T) {
	testWithoutAuth(t)
	time.Sleep(time.Second * 10)
	testMonitor(t)
}

func testWithoutAuth(t *testing.T) {
	todos, err := publicGetTodos.Request()
	assert.Nil(t, err)
	assert.Equal(t, &[]e2e.TodoForResponse{}, todos)
}

func testMonitor(t *testing.T) {
	traces, err := getTraces.RequestWithParam("service=app")
	assert.Nil(t, err)
	assert.NotEmpty(t, traces.Data)
	assert.Len(t, traces.Data[0].Spans, 3)
}
