package controller

import (
	v1 "github.com/gogf/gf-demo-grpc/api/user/v1"
	"github.com/gogf/gf-demo-grpc/internal/logic"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewController)

type Controller struct {
	v1.UnimplementedUserServer

	loc *logic.Logic
}

func NewController(loc *logic.Logic) *Controller {
	return &Controller{
		loc: loc,
	}
}
