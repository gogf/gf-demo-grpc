package server

import (
	v1 "github.com/gogf/gf-demo-grpc/api/user/v1"
	"github.com/gogf/gf-demo-grpc/internal/controller"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

func NewGRPCServer(ctrl *controller.Controller) *grpcx.GrpcServer {
	c := grpcx.Server.NewConfig()
	c.Options = append(c.Options, []grpc.ServerOption{
		grpcx.Server.ChainUnary(
			grpcx.Server.UnaryValidate,
		)}...,
	)
	s := grpcx.Server.New(c)
	// register start
	v1.RegisterUserServer(s.Server, ctrl)
	// register end
	return s
}
