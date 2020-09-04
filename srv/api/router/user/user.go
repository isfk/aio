package user

import (
	"github.com/isfk/aio/srv/api/handler/user"
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

	// 登录注册
	a := e.Group("/v1/user")
	a.POST("/sayHello", user.V1.SayHello)
	a.POST("/login", user.V1.Login)
	a.POST("/logout", user.V1.Logout)
	a.POST("/register", user.V1.Register)

	// 业务接口
	u := e.Group("/v1/user")
	u.Use(middleware.JWT([]byte("secret")))      // auth token
	u.Use(userMiddleware.NeedLogin())            // 需要登录才能访问
	u.Use(casbinMiddleware.Middleware(enforcer)) // 需要有权限才能访问

	u.POST("", user.V1.Add)       // 增
	u.PUT("/:ID", user.V1.Edit)   // 改
	u.DELETE("/:ID", user.V1.Del) // 删

	u.GET("/info", user.V1.Info) // 获取自己信息
	u.GET("/:ID", user.V1.Info)  // 获取其他人信息
	u.GET("/list", user.V1.List) // 获取列表

	u.POST("/role", user.V1.AddRole)                // 添加角色
	u.DELETE("/role/:UID/:RoleID", user.V1.DelRole) // 删除角色
	u.GET("/role/list/:UID", user.V1.RoleList)      // 角色列表

	u.GET("/menu/list", user.V1.MenuList) // 用户菜单列表
}
