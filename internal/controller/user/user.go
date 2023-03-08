package user

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"

	v1 "github.com/gogf/gf-demo-grpc/api/user/v1"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}
