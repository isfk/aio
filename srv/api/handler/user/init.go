package user

import (
	casbinV2 "github.com/casbin/casbin/v2"
	"github.com/isfk/aio/srv/api/pkg/casbin"
	"github.com/isfk/aio/srv/role/proto/role"
	"github.com/isfk/aio/srv/user/proto/user"
	"github.com/micro/go-micro/v2/client"
)

// V1 V1
var V1 *v1

type v1 struct {
	u  user.Service
	r  role.Service
	cr *casbinV2.Enforcer
}

// Init Init
func Init() {

	enforcer, err := casbin.GetEnforcer()
	if err != nil {
		panic(err)
	}

	V1 = &v1{
		u:  user.NewService("go.micro.service.user", client.DefaultClient),
		r:  role.NewService("go.micro.service.role", client.DefaultClient),
		cr: enforcer,
	}
}
