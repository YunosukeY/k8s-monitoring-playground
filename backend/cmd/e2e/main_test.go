package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/YunosukeY/kind-backend/internal/e2e"
	"github.com/stretchr/testify/assert"
)

var login = e2e.NewPostAPI[e2e.User, interface{}]("/api/v1/sessions")
var getTodos = e2e.NewGetAPI[[]e2e.TodoForResponse]("/api/v1/todos")
var postTodo = e2e.NewPostAPI[e2e.TodoForPostRequest, e2e.TodoForResponse]("/api/v1/todos")
var publicGetTodos = e2e.NewGetAPI[[]e2e.TodoForResponse]("/api/v1/public/todos")
var postMail = e2e.NewPostAPI[e2e.Mail, interface{}]("/api/v1/mails")
var getMessages = e2e.NewGetAPI[e2e.Messages]("/api/v2/messages")

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

	expected := e2e.TodoForResponse{ID: 1, Content: "test"}

	todos, err := getTodos.Request()
	assert.Nil(t, err)
	assert.Equal(t, &[]e2e.TodoForResponse{}, todos)

	todo, err := postTodo.Request(e2e.TodoForPostRequest{Content: "test"})
	assert.Nil(t, err)
	assert.Equal(t, &expected, todo)

	todos, err = getTodos.Request()
	assert.Nil(t, err)
	assert.Equal(t, &[]e2e.TodoForResponse{expected}, todos)

	to := "test2@example.com"
	sub := "title"
	msg := "content"
	_, err = postMail.Request(e2e.Mail{To: to, Sub: &sub, Msg: &msg})
	assert.Nil(t, err)
	isOK := false
	for i := 0; i < 10; i++ {
		ms, err := getMessages.Request()
		assert.Nil(t, err)
		if len(ms.Items) != 1 {
			fmt.Println("waiting email")
			time.Sleep(time.Second)
			continue
		}
		isOK = true
		break
	}
	assert.True(t, isOK)
	ms, err := getMessages.Request()
	assert.Nil(t, err)
	expectedMs := &e2e.Messages{
		Items: []e2e.Message{
			{
				Raw: e2e.RawMessage{
					From: "test@example.com",
					To:   []string{to},
					Data: fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, sub, msg),
				},
			},
		},
	}
	assert.Equal(t, expectedMs, ms)
}
