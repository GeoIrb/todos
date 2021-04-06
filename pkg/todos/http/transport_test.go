package http

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/geoirb/todos/pkg/todos"
	tt "github.com/geoirb/todos/pkg/token"
)

const (
	testPort = "8077"
)

func testServer() {
	router := router.New()
	Routing(router, &todos.ServiceMock{}, &tt.TokenMock{})
	httpServer := &fasthttp.Server{
		Handler:          router.Handler,
		DisableKeepalive: true,
	}

	go httpServer.ListenAndServe(":" + testPort)
}
