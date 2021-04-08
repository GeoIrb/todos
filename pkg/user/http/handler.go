package http

import (
	"github.com/valyala/fasthttp"

	"github.com/geoirb/todos/pkg/user"
)

type loginServe struct {
	svc       user.Service
	transport LoginTransport
}

func (s *loginServe) Handler(ctx *fasthttp.RequestCtx) {
	login, err := s.transport.DecodeRequest(&ctx.Request)
	var auth user.Auth
	if err == nil {
		auth, err = s.svc.Login(ctx, login)
	}
	s.transport.EncodeResponse(&ctx.Response, auth, err)
}

func newLoginHandler(svc user.Service, transport LoginTransport) fasthttp.RequestHandler {
	s := loginServe{
		svc:       svc,
		transport: transport,
	}
	return s.Handler
}

type getUserListServe struct {
	svc       user.Service
	transport GetUserListTransport
	token     token
}

func (s *getUserListServe) Handler(ctx *fasthttp.RequestCtx) {
	token, filter, err := s.transport.DecodeRequest(&ctx.Request)
	var users []user.UserInfo
	if err == nil {
		users, err = s.svc.GetUserList(
			s.token.Put(ctx, token),
			filter,
		)
	}
	s.transport.EncodeResponse(&ctx.Response, users, err)
}

func newGetUserListHandler(svc user.Service, transport GetUserListTransport, token token) fasthttp.RequestHandler {
	s := getUserListServe{
		svc:       svc,
		transport: transport,
		token:     token,
	}
	return s.Handler
}

type newUserServe struct {
	svc       user.Service
	transport NewUserTransport
}

func (s *newUserServe) Handler(ctx *fasthttp.RequestCtx) {
	registration, err := s.transport.DecodeRequest(&ctx.Request)
	if err == nil {
		err = s.svc.New(
			ctx,
			registration,
		)
	}
	s.transport.EncodeResponse(&ctx.Response, err)
}

func newNewUserHandler(svc user.Service, transport NewUserTransport) fasthttp.RequestHandler {
	s := newUserServe{
		svc:       svc,
		transport: transport,
	}
	return s.Handler
}

type activateUserServe struct {
	svc       user.Service
	transport ActivateUserTransport
}

func (s *activateUserServe) Handler(ctx *fasthttp.RequestCtx) {
	info, err := s.transport.DecodeRequest(&ctx.Request)
	if err == nil {
		err = s.svc.Create(
			ctx,
			info,
		)
	}
	s.transport.EncodeResponse(&ctx.Response, err)
}

func newActivateUserHandler(svc user.Service, transport ActivateUserTransport) fasthttp.RequestHandler {
	s := activateUserServe{
		svc:       svc,
		transport: transport,
	}
	return s.Handler
}
