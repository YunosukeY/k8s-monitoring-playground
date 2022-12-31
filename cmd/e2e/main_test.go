package main

import (
	"testing"
	"time"

	"github.com/YunosukeY/kind-backend/internal/e2e"
	"github.com/stretchr/testify/assert"
)

var publicGetTodos = e2e.NewGetAPI[[]e2e.TodoForResponse]("/api/v1/public/todos")
var getTraces = e2e.NewGetAPIWithUser[e2e.Traces]("admin", "admin", "localhost:3000", "/api/datasources/proxy/1/api/traces")
var queryData = e2e.NewPostAPIWithUser[e2e.Queries, e2e.QueryResponse]("admin", "admin", "localhost:3000", "/api/ds/query")

func Test(t *testing.T) {
	testWithoutAuth(t)
	time.Sleep(time.Second * 20)
	testMonitor(t)
}

func testWithoutAuth(t *testing.T) {
	todos, err := publicGetTodos.Request()
	assert.Nil(t, err)
	assert.Equal(t, &[]e2e.TodoForResponse{}, todos)
}

func testMonitor(t *testing.T) {
	// test traces
	traces, err := getTraces.RequestWithParam("service=app")
	assert.Nil(t, err)
	assert.NotEmpty(t, traces.Data)
	assert.Len(t, traces.Data[0].Spans, 3)

	// test metrics
	qs := e2e.Queries{
		Queries: []e2e.Query{{
			DSID: 2,
			Expr: "go_memstats_alloc_bytes{job=\"app\"}",
		}},
		From: "now-1m",
		To:   "now",
	}
	qres, err := queryData.Request(qs)
	assert.Nil(t, err)
	assert.NotEmpty(t, qres.Results.A.Frames)
	frame := qres.Results.A.Frames[0]
	assert.Equal(t, "go_memstats_alloc_bytes{instance=\"app.app.svc.cluster.local:8888\", job=\"app\"}", frame.Schema.Name)
	assert.Len(t, frame.Data.Values, 2)
	times := frame.Data.Values[0]
	values := frame.Data.Values[1]
	assert.NotEmpty(t, times)
	assert.NotEmpty(t, values)
	assert.Len(t, times, len(values))
}
