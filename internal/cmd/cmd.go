package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	// Main is the main command.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start grpc server of simple goframe demos",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			grpcx.
			return nil
		},
	}
)
