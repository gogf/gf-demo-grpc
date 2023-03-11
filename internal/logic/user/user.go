package user

import (
	"context"

	"github.com/gogf/gf-demo-grpc/api/pbentity"
	"github.com/gogf/gf-demo-grpc/internal/dao"
	"github.com/gogf/gf-demo-grpc/internal/model/do"
)

type (
	User struct{}
)

func NewLogicUser() *User {
	return &User{}
}

func (s *User) GetById(ctx context.Context, uid uint64) (*pbentity.User, error) {
	var user *pbentity.User
	err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Scan(&user)
	return user, err
}

func (s *User) DeleteById(ctx context.Context, uid uint64) error {
	_, err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Delete()
	return err
}
