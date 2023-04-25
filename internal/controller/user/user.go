package user

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"

	v1 "github.com/gogf/gf-demo-grpc/api/user/v1"
	"github.com/gogf/gf-demo-grpc/internal/dao"
	"github.com/gogf/gf-demo-grpc/internal/model/do"
	"github.com/gogf/gf-demo-grpc/internal/service"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	}).Insert()
	return
}

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

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.User.Ctx(ctx).Page(int(req.Page), int(req.Size)).Scan(&res.Users)
	return
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.User().DeleteById(ctx, req.Id)
	return
}
