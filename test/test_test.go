package test

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"

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
				panic(err)
			}
		}
	})
}
