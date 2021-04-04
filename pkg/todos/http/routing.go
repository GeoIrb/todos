package http

import (
	"context"
	"net/http"

	"github.com/fasthttp/router"

	"github.com/geoirb/todos/pkg/todos"
)

const (
	createTaskURI  = "/task"
	updateTaskURI  = "/task"
	deleteTaskURI  = "/task"
	getTaskListURI = "/task"
)

type token interface {
	Put(ctx context.Context, token string) context.Context
}

func Routing(router *router.Router, svc *todos.Service, token token) {
	router.Handle(http.MethodPost, createTaskURI, newCreateTaskHandler(svc, newCreateTaskTransport(), token))
	router.Handle(http.MethodPatch, updateTaskURI, newUpdateTaskHandler(svc, newUpdateTaskTransport(), token))
	router.Handle(http.MethodDelete, deleteTaskURI, newDeleteTaskHandler(svc, newDeleteTaskTransport(), token))
	router.Handle(http.MethodGet, getTaskListURI, newGetTaskListHandler(svc, newGetTaskListTransport(), token))
}
