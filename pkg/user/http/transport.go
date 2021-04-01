package http

import (
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

func Routing(router *router.Router, svc *user.Service){
	router.Handle(http.MethodPost)
}
