package main

import (
	_ "user-api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"user-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
