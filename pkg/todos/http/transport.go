package http

import (
	"encoding/json"
	"net/http"

	"github.com/geoirb/todos/pkg/todos"
	"github.com/valyala/fasthttp"
)

type CreateTaskTransport interface {
	DecodeRequest(req *fasthttp.Request) (token string, task todos.TaskInfo, err error)
	EncodeResponse(res *fasthttp.Response, err error)
}

type createTaskTransport struct{}

func (t *createTaskTransport) DecodeRequest(req *fasthttp.Request) (token string, task todos.TaskInfo, err error) {
	token = string(req.Header.Peek("Authorization"))
	var request createTaskRequest
	err = json.Unmarshal(req.Body(), &request)
	task = todos.TaskInfo{
		Title:       request.Title,
		Description: request.Description,
		Deadline:    request.Deadline,
	}
	return
}

func (t *createTaskTransport) EncodeResponse(res *fasthttp.Response, err error) {
	responseBuilder(res, nil, err)
}

func newCreateTaskTransport() CreateTaskTransport {
	return &createTaskTransport{}
}

type UpdateTaskTransport interface {
	DecodeRequest(req *fasthttp.Request) (token string, task todos.TaskInfo, err error)
	EncodeResponse(res *fasthttp.Response, err error)
}

type updateTaskTransport struct{}

func (t *updateTaskTransport) DecodeRequest(req *fasthttp.Request) (token string, task todos.TaskInfo, err error) {
	token = string(req.Header.Peek("Authorization"))
	var request updateTaskRequest
	err = json.Unmarshal(req.Body(), &request)
	task = todos.TaskInfo{
		ID:          request.ID,
		Title:       request.Title,
		Description: request.Description,
		Deadline:    request.Deadline,
	}
	return
}
func (t *updateTaskTransport) EncodeResponse(res *fasthttp.Response, err error) {
	responseBuilder(res, nil, err)
}

func newUpdateTaskTransport() UpdateTaskTransport {
	return &updateTaskTransport{}
}

type DeleteTaskTransport interface {
	DecodeRequest(req *fasthttp.Request) (token string, filter todos.Filter, err error)
	EncodeResponse(res *fasthttp.Response, err error)
}

type deleteTaskTransport struct{}

func (t *deleteTaskTransport) DecodeRequest(req *fasthttp.Request) (token string, filter todos.Filter, err error) {
	token = string(req.Header.Peek("Authorization"))
	var request deleteTaskRequest
	err = json.Unmarshal(req.Body(), &request)
	filter = todos.Filter{
		ID: &request.ID,
	}
	return
}
func (t *deleteTaskTransport) EncodeResponse(res *fasthttp.Response, err error) {
	responseBuilder(res, nil, err)
}

func newDeleteTaskTransport() DeleteTaskTransport {
	return &deleteTaskTransport{}
}

type GetTaskListTransport interface {
	DecodeRequest(req *fasthttp.Request) (token string, filter todos.Filter, err error)
	EncodeResponse(res *fasthttp.Response, tasks []todos.TaskInfo, err error)
}

type getTaskListTransport struct{}

func (t *getTaskListTransport) DecodeRequest(req *fasthttp.Request) (token string, filter todos.Filter, err error) {
	token = string(req.Header.Peek("Authorization"))
	var request deleteTaskRequest
	err = json.Unmarshal(req.Body(), &request)
	filter = todos.Filter{
		ID: &request.ID,
	}
	return
}
func (t *getTaskListTransport) EncodeResponse(res *fasthttp.Response, tasks []todos.TaskInfo, err error) {
	var response getTaskListResponse
	response.Tasks = make([]taskInfo, 0, len(tasks))
	for _, task := range tasks {
		response.Tasks = append(response.Tasks, taskInfo(task))
	}
	responseBuilder(res, response, err)
}

func newGetTaskListTransport() GetTaskListTransport {
	return &getTaskListTransport{}
}

var responseCode map[error]int = map[error]int{
	todos.ErrFailedAuthenticate: http.StatusUnauthorized,
}

func responseBuilder(res *fasthttp.Response, response interface{}, err error) {
	code := http.StatusOK
	if err != nil {
		var isExist bool
		if code, isExist = responseCode[err]; !isExist {
			code = http.StatusInternalServerError
		}
		res.SetStatusCode(code)
		res.SetBody([]byte(err.Error()))
		return
	}

	body, errMarshal := json.Marshal(response)
	if errMarshal != nil {
		code = http.StatusInternalServerError
		body = []byte(errMarshal.Error())
	}
	res.SetStatusCode(code)
	res.SetBody(body)
}
