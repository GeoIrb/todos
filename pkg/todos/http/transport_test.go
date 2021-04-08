package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/fasthttp/router"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"github.com/geoirb/todos/pkg/todos"
	tt "github.com/geoirb/todos/pkg/token"
)

const (
	testPort = "7070"

	testToken       = "test-token"
	testTitle       = "test-title"
	testDescription = "test-description"
	testDeadline    = 123

	testErrorMessage = "test-error-message"

	successCreateTaskAnswer = "null"
)

var (
	nilError  error
	testError = errors.New(testErrorMessage)
	testTask  = todos.TaskInfo{
		Title:       testTitle,
		Description: testDescription,
		Deadline:    testDeadline,
	}
)

func TestCreateTask(t *testing.T) {
	var testCreateTaskList = []struct {
		name           string
		svcReturn      interface{}
		expectedCode   int
		expectedAnswer string
	}{
		{
			"Success Test",
			nilError,
			http.StatusOK,
			successCreateTaskAnswer,
		},
		{
			"Failed Authenticate Test",
			todos.ErrFailedAuthenticate,
			http.StatusUnauthorized,
			todos.ErrFailedAuthenticate.Error(),
		},
		{
			"Failed Test",
			testError,
			http.StatusInternalServerError,
			testErrorMessage,
		},
	}

	tokenMock := &tt.TokenMock{}
	tokenMock.On("Put", testToken).Return(nil)

	for _, test := range testCreateTaskList {
		t.Run(test.name, func(t *testing.T) {
			svcMock := &todos.ServiceMock{}
			svcMock.On("CreateTask", testTask).
				Return(test.svcReturn)

			server := startTestServer(t, svcMock, tokenMock)
			defer server.Shutdown()

			taskRequest := createTaskRequest{
				Title:       testTask.Title,
				Description: testTask.Description,
				Deadline:    testTask.Deadline,
			}
			data, err := json.Marshal(taskRequest)
			assert.NoError(t, err, "marshal request")
			code, answer := sendRequest(t, "POST", createTaskURI, data)
			assert.Equal(t, test.expectedCode, code, "code request")
			assert.Equal(t, test.expectedAnswer, answer, "body request")
		})
	}
}

func sendRequest(t *testing.T, method, uri string, body []byte) (code int, answer string) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI("http://127.0.0.1:" + testPort + uri)

	req.Header.SetMethod(method)
	req.Header.Set("Authorization", testToken)
	req.SetBody(body)

	client := &fasthttp.HostClient{
		Addr: "127.0.0.1:" + testPort,
	}
	err := client.Do(req, resp)
	assert.NoError(t, err, "send request")

	return resp.StatusCode(), string(resp.Body())
}

func startTestServer(t *testing.T, svc todos.Service, token token) *fasthttp.Server {
	router := router.New()
	Routing(router, svc, token)
	httpServer := &fasthttp.Server{
		Handler:          router.Handler,
		DisableKeepalive: true,
	}

	go httpServer.ListenAndServe(":" + testPort)

	time.Sleep(time.Second)

	return httpServer
}
