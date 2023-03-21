//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gogf/gf-demo-grpc/internal/controller"
	"github.com/gogf/gf-demo-grpc/internal/logic"
	"github.com/gogf/gf-demo-grpc/internal/server"
	"github.com/google/wire"
)

// initApp di kratos application.
func initApp() *cmdApp {
	panic(wire.Build(
		logic.ProviderSet,
		controller.ProviderSet,
		server.ProviderSet,
		newApp,
	))
}
