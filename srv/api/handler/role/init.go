package role

import (
	casbinV2 "github.com/casbin/casbin/v2"
	"github.com/isfk/aio/proto/role"
	"github.com/isfk/aio/srv/api/pkg/casbin"
	"github.com/micro/go-micro/v2/client"
)

// V1 V1
var V1 *v1

type v1 struct {
	r role.Service
	e *casbinV2.Enforcer
}

// Init Init
func Init() {
	enforcer, err := casbin.GetEnforcer()
	if err != nil {
		panic(err)
	}

	V1 = &v1{
		r: role.NewService("go.micro.service.role", client.DefaultClient),
		e: enforcer,
	}
}
