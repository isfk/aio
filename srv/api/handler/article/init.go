package article

import (
	"github.com/isfk/aio/srv/article/proto/article"
	"github.com/micro/go-micro/v2/client"
)

// V1 V1
var V1 *v1

type v1 struct {
	a article.Service
}

// Init Init
func Init() {
	V1 = &v1{
		a: article.NewService("go.micro.service.article", client.DefaultClient),
	}
}
