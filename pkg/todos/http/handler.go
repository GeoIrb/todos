package http

import (
	"github.com/valyala/fasthttp"

	"github.com/geoirb/todos/pkg/todos"
)

type createTaskServe struct {
	svc       *todos.Service
	transport CreateTaskTransport
	token     token
}

func (s *createTaskServe) Handler(ctx *fasthttp.RequestCtx) {
	token, task, err := s.transport.DecodeRequest(&ctx.Request)
	if err == nil {
		err = s.svc.CreateTask(
			s.token.Put(ctx, token),
			task,
		)
	}
	s.transport.EncodeResponse(&ctx.Response, err)
}

func newCreateTaskHandler(svc *todos.Service, transport CreateTaskTransport, token token) fasthttp.RequestHandler {
	s := &createTaskServe{
		svc:       svc,
		transport: transport,
		token:     token,
	}
	return s.Handler
}

type updateTaskServe struct {
	svc       *todos.Service
	transport UpdateTaskTransport
	token     token
}

func (s *updateTaskServe) Handler(ctx *fasthttp.RequestCtx) {
	token, task, err := s.transport.DecodeRequest(&ctx.Request)
	if err == nil {
		err = s.svc.UpdateTask(
			s.token.Put(ctx, token),
			task,
		)
	}
	s.transport.EncodeResponse(&ctx.Response, err)
}

func newUpdateTaskHandler(svc *todos.Service, transport UpdateTaskTransport, token token) fasthttp.RequestHandler {
	s := &updateTaskServe{
		svc:       svc,
		transport: transport,
		token:     token,
	}
	return s.Handler
}

type deleteTaskServe struct {
	svc       *todos.Service
	transport DeleteTaskTransport
	token     token
}

func (s *deleteTaskServe) Handler(ctx *fasthttp.RequestCtx) {
	token, filter, err := s.transport.DecodeRequest(&ctx.Request)
	if err == nil {
		err = s.svc.DeleteTask(
			s.token.Put(ctx, token),
			filter,
		)
	}
	s.transport.EncodeResponse(&ctx.Response, err)
}

func newDeleteTaskHandler(svc *todos.Service, transport DeleteTaskTransport, token token) fasthttp.RequestHandler {
	s := &deleteTaskServe{
		svc:       svc,
		transport: transport,
		token:     token,
	}
	return s.Handler
}

type getTaskListServe struct {
	svc       *todos.Service
	transport GetTaskListTransport
	token     token
}

func (s *getTaskListServe) Handler(ctx *fasthttp.RequestCtx) {
	token, filter, err := s.transport.DecodeRequest(&ctx.Request)
	var tasks []todos.TaskInfo
	if err == nil {
		tasks, err = s.svc.GetTaskList(
			s.token.Put(ctx, token),
			filter,
		)
	}
	s.transport.EncodeResponse(&ctx.Response, tasks, err)
}

func newGetTaskListHandler(svc *todos.Service, transport GetTaskListTransport, token token) fasthttp.RequestHandler {
	s := &getTaskListServe{
		svc:       svc,
		transport: transport,
		token:     token,
	}
	return s.Handler
}
