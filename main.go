package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf-demo-grpc/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf-demo-grpc/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
