package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta   `path:"/user" tags:"User" method:"put" summary:"User registry"`
	Passport string `v:"required"`
	Password string `v:"required"`
	Nickname string `v:"required"`
}
type RegisterRes struct {
}
