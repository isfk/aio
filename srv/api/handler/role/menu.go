package role

import (
	"context"
	"fmt"
	"github.com/isfk/aio/proto/role"
	rolePkg "github.com/isfk/aio/srv/api/pkg/role"
	"net/http"
	"strconv"

	casbinV2 "github.com/casbin/casbin/v2"
	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/isfk/aio/srv/api/pkg/api"
	"github.com/isfk/aio/srv/api/pkg/validate"

	"github.com/labstack/echo/v4"
)

// MenuList MenuList
func (v1 *v1) MenuList(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	menuList, err := v1.r.MenuList(context.Background(), &role.Null{})
	if err != nil {
		log.Error("v1.r.MenuList err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data: 	 menuList.List,
		Message: "获取成功",
	})
}

// AddRoleMenu 添加角色菜单
func (v1 *v1) AddRoleMenu(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args AddRoleMenuArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.casbin.AddRoleMenu args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	roleIDString := strconv.FormatInt(args.RoleID, 10)

	v1.e.RemoveFilteredPolicy(0, fmt.Sprintf("role_id:%s.menu", roleIDString))

	menuList, err := v1.r.MenuList(context.Background(), &role.Null{})
	if err != nil {
		log.Error("v1.r.InfoByUsername err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	menuListAll := rolePkg.GetMenuListAll(menuList.List)

	go func(e *casbinV2.Enforcer, IDList []int64) {
		for _, v := range args.IDList {
			fmt.Println(v)
			if a, ok := menuListAll[v]; ok {
				fmt.Println(a)
				ok, err := v1.e.AddPolicy(fmt.Sprintf("role_id:%s.menu", roleIDString), a.Path, a.Path, strconv.FormatInt(v, 10))

				if err != nil {
					log.Error("v1.e.AddPolicy err ", err.Error())
				}

				if ok == false {
					log.Error("v1.e.AddPolicy err ", "添加失败")
				}
			}
		}
	}(v1.e, args.IDList)

	return c.JSON(http.StatusOK, &api.RetOk{
		Message: "添加成功",
	})
}

// RoleMenus 角色的菜单
func (v1 *v1) RoleMenus(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args RulesArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.AddRole args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    map[string]interface{}{"list": v1.e.GetFilteredPolicy(0, fmt.Sprintf("role_id:%d.menu", args.RoleID))},
		Message: "获取成功",
	})
}
