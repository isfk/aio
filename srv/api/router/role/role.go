package role

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/isfk/aio/srv/api/handler/role"
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
	r := e.Group("/v1/role")
	r.Use(middleware.JWT([]byte("secret")))      // auth token
	r.Use(userMiddleware.NeedLogin())            // 需要登录才能访问
	r.Use(casbinMiddleware.Middleware(enforcer)) // 需要有权限才能访问

	r.POST("", role.V1.Add)       // 增
	r.PUT("/:ID", role.V1.Edit)   // 改
	r.DELETE("/:ID", role.V1.Del) // 删

	r.GET("/:ID", role.V1.Info)  // 获取信息
	r.GET("/list", role.V1.List) // 获取列表

	r.GET("/rule/list", role.V1.RuleList)      // 规则列表
	r.POST("/rules", role.V1.AddRoleRule)      // 给角色添加规则
	r.GET("/rules/:RoleID", role.V1.RoleRules) // 角色的规则

	r.GET("/menu/list", role.V1.MenuList)      // 菜单列表
	r.POST("/menus", role.V1.AddRoleMenu)      // 给角色添加菜单
	r.GET("/menus/:RoleID", role.V1.RoleMenus) // 角色的菜单
}
