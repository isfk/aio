package article

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/isfk/aio/srv/api/handler/article"
	casbinMiddleware "github.com/isfk/aio/srv/api/middleware/casbin"
	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	"github.com/isfk/aio/srv/api/pkg/casbin"
)

// Register Register
func Register(e *echo.Echo) {
	enforcer, err := casbin.GetEnforcer()
	if err != nil {
		panic(err)
	}
	enforcer.LoadPolicy()

	// 业务接口
	u := e.Group("/v1/article")
	u.Use(middleware.JWT([]byte("secret")))      // auth token
	u.Use(userMiddleware.NeedLogin())            // 需要登录才能访问
	u.Use(casbinMiddleware.Middleware(enforcer)) // 需要有权限才能访问

	u.POST("", article.V1.Add)       // 增
	u.PUT("/:ID", article.V1.Edit)   // 改
	u.DELETE("/:ID", article.V1.Del) // 删

	u.GET("/:ID", article.V1.Info)  // 获取信息
	u.GET("/list", article.V1.List) // 获取列表
}
