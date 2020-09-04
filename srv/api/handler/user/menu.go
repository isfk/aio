package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/isfk/aio/srv/api/pkg/api"
	log "github.com/micro/go-micro/v2/logger"

	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	"github.com/isfk/aio/srv/api/pkg/validate"
	roleProto "github.com/isfk/aio/srv/role/proto/role"
	"github.com/labstack/echo/v4"
)

// MenuList MenuList
func (v1 *v1) MenuList(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args EmptyArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.MenuList args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 拿到全部菜单，然后进行权限验证即可
	menuList, err := v1.r.MenuList(context.Background(), &roleProto.Null{})
	if err != nil {
		log.Error("v1.r.MenuList err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	var newMenuList []*roleProto.Menu

	for _, v := range menuList.List {
		menu := &roleProto.Menu{}

		// 父级权限
		ok, err := v1.cr.Enforce(fmt.Sprintf("uid:%d", userInfo.Id), v.Path, v.Path, strconv.FormatInt(v.Id, 10))
		if err != nil {
			log.Error("v1.cr.Enforce err ", err.Error())
			return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
		}

		if ok == false && userInfo.Id != 1 {
			continue
		}

		var children []*roleProto.Menu

		if len(v.Children) > 0 {
			for n, m := range v.Children {
				// 下级权限
				if userInfo.Id == 1 {
					children = append(children, m)
					continue
				}

				ok, err := v1.cr.Enforce(fmt.Sprintf("uid:%d", userInfo.Id), m.Path, m.Path, strconv.FormatInt(m.Id, 10))
				if err != nil {
					log.Error("v1.cr.Enforce err ", err.Error())
					return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
				}

				if ok == false {
					m = &roleProto.Menu{}
					v.Children[n] = &roleProto.Menu{}
					continue
				}

				children = append(children, m)
			}
		}

		menu = v
		if len(children) > 0 {
			menu.Children = children
		}
		newMenuList = append(newMenuList, menu)
	}

	return c.JSON(http.StatusOK, &api.RetOk{Data: newMenuList, Message: "获取成功"})
}
