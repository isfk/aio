package router

import (
	"github.com/isfk/aio/srv/api/handler"
	articleRouter "github.com/isfk/aio/srv/api/router/article"
	roleRouter "github.com/isfk/aio/srv/api/router/role"
	userRouter "github.com/isfk/aio/srv/api/router/user"
	"github.com/labstack/echo/v4"
)

// Router Router
func Router(e *echo.Echo) *echo.Echo {
	handler.Init()

	e.File("/", "html/index.html")

	articleRouter.Register(e)
	userRouter.Register(e)
	roleRouter.Register(e)
	return e
}
