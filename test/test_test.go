package test

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"

	v1 "github.com/gogf/gf-demo-grpc/api/user/v1"
)

func Test_Create(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("demo")
			user = v1.NewUserClient(conn)
		)
		for i := 1; i <= 10; i++ {
			_, err := user.Create(ctx, &v1.CreateReq{
				Passport: fmt.Sprintf(`passport-%d`, i),
				Password: "123456",
				Nickname: fmt.Sprintf(`nickname-%d`, i),
			})
			if err != nil {
				g.Log().Fatalf(ctx, `create user failed: %+v`, err)
			}
		}
	})
}

func Test_GetList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx   = gctx.GetInitCtx()
			conn  = grpcx.Client.MustNewGrpcClientConn("demo")
			user  = v1.NewUserClient(conn)
			res   *v1.GetListRes
			err   error
			size  int32 = 2
			pages       = []int32{1, 2, 3}
		)
		for _, page := range pages {
			res, err = user.GetList(ctx, &v1.GetListReq{
				Page: page,
				Size: size,
			})
			if err != nil {
				g.Log().Fatalf(ctx, `get user list failed: %+v`, err)
			}
			fmt.Println(res.Users)
		}
	})
}

func Test_GetOne(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("demo")
			user = v1.NewUserClient(conn)
			res  *v1.GetOneRes
			err  error
		)
		res, err = user.GetOne(ctx, &v1.GetOneReq{
			Id: 1,
		})
		if err != nil {
			g.Log().Fatalf(ctx, `get user failed: %+v`, err)
		}
		fmt.Println(res.User)

		res, err = user.GetOne(ctx, &v1.GetOneReq{
			Id: 100,
		})
		if err != nil {
			g.Log().Fatalf(ctx, `get user failed: %+v`, err)
		}
		fmt.Println(res.User)
	})
}

func Test_Delete(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("demo")
			user = v1.NewUserClient(conn)
		)
		_, err := user.Delete(ctx, &v1.DeleteReq{
			Id: 1,
		})
		if err != nil {
			g.Log().Fatalf(ctx, `get user failed: %+v`, err)
		}

		getOneRes, err := user.GetOne(ctx, &v1.GetOneReq{
			Id: 1,
		})
		if err != nil {
			g.Log().Fatalf(ctx, `delete user failed: %+v`, err)
		}
		fmt.Println(getOneRes.User)
	})
}

func Test_Validation(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("demo")
			user = v1.NewUserClient(conn)
		)
		_, err := user.Delete(ctx, &v1.DeleteReq{})
		if err != nil {
			g.Log().Fatalf(ctx, `delete user failed: %+v`, err)
		}
	})
}
