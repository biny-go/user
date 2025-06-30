package main

import (
	api "user/api"
	c "user/controller"
	r "user/repo"
	s "user/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

var ProviderControllerSet = wire.NewSet(c.NewHelloController)
var ProviderServiceSet = wire.NewSet(s.NewHelloService)
var ProviderRepoSet = wire.NewSet(r.NewHelloRepo)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(helloController *c.HelloController) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	srv := http.NewServer(opts...)

	api.RegisterHelloHTTPServer(srv, helloController)

	return srv
}

var ProviderServerSet = wire.NewSet(NewHTTPServer)

// wireApp init kratos application.
func wireApp() (*kratos.App, func(), error) {
	panic(wire.Build(ProviderControllerSet, ProviderServiceSet, ProviderRepoSet, ProviderServerSet, newApp))
}
