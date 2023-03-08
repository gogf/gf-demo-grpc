package user

import (
	"context"

	"github.com/gogf/gf-demo-grpc/internal/dao"
	"github.com/gogf/gf-demo-grpc/internal/model/do"
)

type (
	sUser struct{}
)

func init() {

}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) DeleteById(ctx context.Context, uid uint) error {
	_, err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Delete()
	return err
}
