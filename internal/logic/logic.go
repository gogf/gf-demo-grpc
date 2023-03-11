package logic

import (
	"github.com/gogf/gf-demo-grpc/internal/logic/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewLogic, user.NewLogicUser)

type Logic struct {
	*user.User
}

func NewLogic(user *user.User) *Logic {
	return &Logic{
		User: user,
	}
}
