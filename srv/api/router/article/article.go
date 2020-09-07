package article

import (
	"github.com/isfk/aio/srv/api/handler/article"
	casbinMiddleware "github.com/isfk/aio/srv/api/middleware/casbin"
	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	"github.com/isfk/aio/srv/api/pkg/casbin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Register Register
func Register(e *echo.Echo) {
	enforcer, err := casbin.GetEnforcer()
	if err != nil {
		panic(err)
	}
	enforcer.LoadPolicy()

	// 业务接口
	a := e.Group("/v1/article")
	a.Use(middleware.JWT([]byte("secret")))      // auth token
	a.Use(userMiddleware.NeedLogin())            // 需要登录才能访问
	a.Use(casbinMiddleware.Middleware(enforcer)) // 需要有权限才能访问

	a.POST("", article.V1.Add)       // 增
	a.PUT("/:ID", article.V1.Edit)   // 改
	a.DELETE("/:ID", article.V1.Del) // 删

	a.GET("/:ID", article.V1.Info)  // 获取信息
	a.GET("/list", article.V1.List) // 获取列表

	a.POST("", article.V1.AddCategory)       // 增
	a.PUT("/:ID", article.V1.EditCategory)   // 改
	a.DELETE("/:ID", article.V1.DelCategory) // 删

	a.GET("/:ID", article.V1.CategoryInfo)  // 获取信息
	a.GET("/list", article.V1.CategoryList) // 获取列表

}
