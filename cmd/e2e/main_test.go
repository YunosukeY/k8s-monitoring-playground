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
	time.Sleep(time.Second * 30)
	testMonitor(t)
}

func testWithoutAuth(t *testing.T) {
	todos, err := publicGetTodos.Request()
	assert.Nil(t, err)
	assert.Equal(t, &[]e2e.TodoForResponse{}, todos)
}

// test metrics are retrieved via Grafana
func testMonitor(t *testing.T) {
	// test traces
	traces, err := getTraces.RequestWithParam("service=app")
	assert.Nil(t, err)
	assert.NotEmpty(t, traces.Data)
	assert.Len(t, traces.Data[0].Spans, 4)

	// test metrics
	metricsQ := e2e.Queries{
		Queries: []e2e.Query{{
			DSID: 2,
			Expr: "go_memstats_alloc_bytes{job=\"app\"}",
		}},
		From: "now-1m",
		To:   "now",
	}
	qres, err := queryData.Request(metricsQ)
	assert.Nil(t, err)
	assert.Equal(t, 200, qres.Results.A.Status)
	assert.NotEmpty(t, qres.Results.A.Frames)

	// test logs
	logsQ := e2e.Queries{
		Queries: []e2e.Query{{
			DSID: 3,
			Expr: "{job=\"app/app\"} |= ``",
		}},
		From: "now-1m",
		To:   "now",
	}
	qres, err = queryData.Request(logsQ)
	assert.Nil(t, err)
	assert.Equal(t, 200, qres.Results.A.Status)
	assert.NotEmpty(t, qres.Results.A.Frames)
}
