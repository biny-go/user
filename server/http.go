package server

import (
	"context"
	"fmt"
	c "user/controller"

	"github.com/biny-go/loong/server"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(helloController *c.HelloController) *http.Server {

	registrars := []server.HTTPServiceRegistrar{
		helloController,
	}

	options := server.HTTPServerOptions{
		Address: ":8081",
		Middlewares: []middleware.Middleware{
			demoMiddleware(),
		},
		Registrars: registrars,
	}
	srv := server.NewHTTPServer(options)

	// for _, r := range registrars {
	// 	r.Register(srv)
	// }

	return srv
}

// demoMiddleware 是一个简单的中间件
func demoMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if ts, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := ts.(http.Transporter); ok {
					fmt.Printf("Request received: %s %s\n", ht.Request().Method, ht.Request().URL.Path)
				}
			}

			resp, err := handler(ctx, req)

			if ts, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := ts.(http.Transporter); ok {
					fmt.Printf("Request after: %s %s\n", ht.Request().Method, ht.Request().URL.Path)
				}
			}

			return resp, err
		}
	}
}
