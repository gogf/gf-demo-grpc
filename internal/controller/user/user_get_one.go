package user

import (
	"context"

	v1 "github.com/gogf/gf-demo-grpc/api/user/v1"
	"github.com/gogf/gf-demo-grpc/internal/service"
)

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	user, err := service.User().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetOneRes{
		User: user,
	}
	return
}
