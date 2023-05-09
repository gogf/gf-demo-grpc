package user

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"

	v1 "github.com/gogf/gf-demo-grpc/app/user-api/api/user/v1"
	userV1 "github.com/gogf/gf-demo-grpc/app/user-rpc/api/user/v1"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Hello(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	client := userV1.NewUserClient(grpcx.Client.MustNewGrpcClientConn("user"))
	_, err = client.Create(ctx, &userV1.CreateReq{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}
