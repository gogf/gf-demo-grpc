package user

import (
	"context"

	"github.com/gogf/gf-demo-grpc/app/user-rpc/api/pbentity"
	"github.com/gogf/gf-demo-grpc/app/user-rpc/internal/dao"
	"github.com/gogf/gf-demo-grpc/app/user-rpc/internal/model/do"
	"github.com/gogf/gf-demo-grpc/app/user-rpc/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(&sUser{})
}

func (s *sUser) GetById(ctx context.Context, uid uint64) (*pbentity.User, error) {
	var user *pbentity.User
	err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Scan(&user)
	return user, err
}

func (s *sUser) DeleteById(ctx context.Context, uid uint64) error {
	_, err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Delete()
	return err
}
