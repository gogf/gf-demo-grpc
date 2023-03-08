package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"google.golang.org/grpc"

	v1 "github.com/gogf/gf-demo-grpc/api/user/v1"
)

func (*Controller) Create(ctx context.Context, in *v1.CreateReq, opts ...grpc.CallOption) (*v1.CreateRes, error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetOne(ctx context.Context, in *v1.GetOneReq, opts ...grpc.CallOption) (*v1.GetOneRes, error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetList(ctx context.Context, in *v1.GetListReq, opts ...grpc.CallOption) (*v1.GetListRes, error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Delete(ctx context.Context, in *v1.DeleteReq, opts ...grpc.CallOption) (*v1.DeleteRes, error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
