package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"

	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/gogf/gf-demo-grpc/internal/controller/user"
)

var (
	// Main is the main command.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start grpc server of simple goframe demos",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := grpcx.Server.New(&grpcx.GrpcServerConfig{
				Name: "demo",
			})
			user.Register(s)
			s.Run()
			return nil
		},
	}
)
