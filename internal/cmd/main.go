package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/net/ghttp"
	"sync"

	_ "github.com/gogf/gf-demo-grpc/internal/logic"
)

func main() {
	initApp()
}

func (c *cmdApp) runApp() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		c.Server.Run()
	}()
	go func() {
		c.GrpcServer.Run()
	}()
	wg.Wait()
}

type cmdApp struct {
	*grpcx.GrpcServer
	*ghttp.Server
}

func newApp(r *grpcx.GrpcServer, h *ghttp.Server) *cmdApp {
	return &cmdApp{
		GrpcServer: r,
		Server:     h,
	}
}
