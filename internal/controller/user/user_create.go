package user

import (
	"context"

	v1 "github.com/gogf/gf-demo-grpc/api/user/v1"
	"github.com/gogf/gf-demo-grpc/internal/dao"
	"github.com/gogf/gf-demo-grpc/internal/model/do"
)

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	res = &v1.CreateRes{}
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	}).Insert()
	return
}
