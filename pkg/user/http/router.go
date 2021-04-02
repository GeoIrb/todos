package http

import (
	"context"
	"net/http"

	"github.com/fasthttp/router"
	"github.com/geoirb/todos/pkg/user"
)

const (
	loginURI            = "/login"
	getUserList         = "/user"
	registrationUserURI = "/user/registration"
	createUserURI       = "/user/create"
)

type token interface {
	Put(ctx context.Context, token string) context.Context
}

func Routing(router *router.Router, svc *user.Service, token token) {
	router.Handle(http.MethodPost, loginURI, newLoginHandler(svc, newLoginTransport()))
	router.Handle(http.MethodGet, getUserList, newGetUserListHandler(svc, newGetUserTransport(), token))
	router.Handle(http.MethodPost, registrationUserURI, newRegistrationUserHandler(svc, NewRegistrationUserTransport()))
	router.Handle(http.MethodPut, createUserURI, newCreateUserHandler(svc, NewCreateUserTransport()))
}
