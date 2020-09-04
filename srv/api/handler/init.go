package handler

import (
	"github.com/isfk/aio/srv/api/handler/article"
	"github.com/isfk/aio/srv/api/handler/role"
	"github.com/isfk/aio/srv/api/handler/user"
)

// Init Init
func Init() {
	article.Init()
	role.Init()
	user.Init()
}
