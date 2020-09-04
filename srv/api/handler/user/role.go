package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/isfk/aio/srv/api/pkg/api"

	roleProto "github.com/isfk/aio/proto/role"
	proto "github.com/isfk/aio/proto/user"
	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	"github.com/isfk/aio/srv/api/pkg/validate"
	"github.com/labstack/echo/v4"
	log "github.com/micro/go-micro/v2/logger"
)

// AddRole AddRole
func (v1 *v1) AddRole(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args AddRoleArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Infof("v1.AddRole args %v", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	info := &proto.UserRole{
		UserId: args.UserID,
		RoleId: args.RoleID,
	}

	oldUserRole, err := v1.u.UserRoleInfo(context.Background(), info)
	if err != nil {
		log.Error("v1.u.UserRoleInfo err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if oldUserRole.Id > 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户角色已存在"})
	}

	resp, err := v1.u.CreateUserRole(context.Background(), info)

	if err != nil {
		log.Error("v1.u.CreateUserRole err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	ok, err := v1.cr.AddGroupingPolicy(fmt.Sprintf("user_id:%d", args.UserID), fmt.Sprintf("role_id:%d.api", args.RoleID))

	if err != nil {
		log.Error("v1.cr.AddGroupingPolicy err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if ok == false {
		log.Error("v1.cr.AddGroupingPolicy err", "添加接口权限失败")
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "添加接口权限失败"})
	}

	ok, err = v1.cr.AddGroupingPolicy(fmt.Sprintf("user_id:%d", args.UserID), fmt.Sprintf("role_id:%d.menu", args.RoleID))

	if err != nil {
		log.Error("v1.cr.AddGroupingPolicy err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if ok == false {
		log.Error("v1.cr.AddGroupingPolicy err", "添加菜单权限失败")
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "添加菜单权限失败"})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "添加成功",
	})
}

// DelRole DelRole
func (v1 *v1) DelRole(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args DelRoleArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Infof("v1.DelRole args %v", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	_, err := v1.u.DeleteUserRole(context.Background(), &proto.UserRole{
		UserId: args.UserID,
		RoleId: args.RoleID,
	})

	if err != nil {
		log.Error("v1.u.DeleteUserRole err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	ok, err := v1.cr.RemoveGroupingPolicy(fmt.Sprintf("user_id:%d", args.UserID), fmt.Sprintf("role_id:%d.api", args.RoleID))

	if err != nil {
		log.Error("v1.cr.RemoveGroupingPolicy err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if ok == false {
		log.Error("v1.cr.RemoveGroupingPolicy err", "删除接口权限失败")
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "删除接口权限失败"})
	}

	ok, err = v1.cr.RemoveGroupingPolicy(fmt.Sprintf("user_id:%d", args.UserID), fmt.Sprintf("role_id:%d.menu", args.RoleID))

	if err != nil {
		log.Error("v1.cr.RemoveGroupingPolicy err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if ok == false {
		log.Error("v1.cr.RemoveGroupingPolicy err", "删除菜单权限失败")
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "删除菜单权限失败"})
	}

	return c.JSON(http.StatusOK, &api.RetOk{Message: "删除成功"})
}

// RoleList RoleList
func (v1 *v1) RoleList(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args RoleListArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Infof("v1.RoleList args %v", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	resp, err := v1.u.UserRoleListByUserID(context.Background(), &proto.UserID{
		UserId: args.UserID,
	})

	if err != nil {
		log.Error("v1.u.UserRoleListByUID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	for _, v := range resp.List {
		v.Role, err = v1.r.InfoByID(context.Background(), &roleProto.ID{Id: v.RoleId})
		if err != nil {
			return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "获取成功",
	})
}
