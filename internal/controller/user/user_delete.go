package user

import (
	"context"

	v1 "github.com/gogf/gf-demo-grpc/api/user/v1"
	"github.com/gogf/gf-demo-grpc/internal/service"
)

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.User().DeleteById(ctx, req.Id)
	return
}
