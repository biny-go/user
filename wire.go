//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	c "user/controller"
	r "user/repo"
	serve "user/server"
	s "user/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// ProviderSet 定义所有依赖
var ProviderSet = wire.NewSet(
	r.NewHelloRepo,
	s.NewHelloService,
	c.NewHelloController,
	serve.NewHTTPServer,
	newApp,
)

// wireApp 初始化 Kratos App
func wireApp() (*kratos.App, func(), error) {
	wire.Build(ProviderSet) // [!code highlight]
	return &kratos.App{}, nil, nil
}

// wireApp init kratos application.
// func wireApp() (*kratos.App, func(), error) {
// 	panic(wire.Build(c.NewHelloController, s.NewHelloService, r.NewHelloRepo, server.NewHTTPServer, newApp))
// }
