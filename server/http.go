package server

import (
	api "user/api"
	c "user/controller"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

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
