package handler

import (
	"context"
	"net/http"

	"github.com/isfk/aio/srv/role/proto/role"
)

const (
	// MethodGet get
	MethodGet = http.MethodGet

	// MethodPos post
	MethodPos = http.MethodPost

	// MethodPut put
	MethodPut = http.MethodPut

	// MethodDel delete
	MethodDel = http.MethodDelete
)

// RuleList RuleList
func (e *Role) RuleList(ctx context.Context, req *role.Null, rsp *role.RuleListRet) error {
	var list []*role.Api
	list = append(list, &role.Api{
		Id:    10000,
		Label: "用户",
		Children: []*role.Api{
			{Id: 10001, Method: MethodPos, Label: "登录", Path: "/api/v1/user/login"},
			{Id: 10002, Method: MethodPos, Label: "注册", Path: "/api/v1/user/register"},
			{Id: 10003, Method: MethodPos, Label: "退出", Path: "/api/v1/user/logout"},
			{Id: 10004, Method: MethodGet, Label: "个人信息", Path: "/api/v1/user/info"},
			{Id: 10005, Method: MethodGet, Label: "他人信息", Path: "/api/v1/user/:id"},
			{Id: 10006, Method: MethodPos, Label: "添加用户", Path: "/api/v1/user"},
			{Id: 10007, Method: MethodPut, Label: "编辑用户", Path: "/api/v1/user/:id"},
			{Id: 10008, Method: MethodDel, Label: "删除用户", Path: "/api/v1/user/:id"},
			{Id: 10009, Method: MethodGet, Label: "用户列表", Path: "/api/v1/user/list"},
			{Id: 10010, Method: MethodGet, Label: "菜单列表", Path: "/api/v1/user/menu/list"},
		},
	})

	list = append(list, &role.Api{
		Id:    10100,
		Label: "用户角色",
		Children: []*role.Api{
			{Id: 10101, Method: MethodPos, Label: "添加角色", Path: "/api/v1/user/role"},
			{Id: 10102, Method: MethodDel, Label: "删除角色", Path: "/api/v1/user/role/:uid/:role_id"},
			{Id: 10103, Method: MethodGet, Label: "角色列表", Path: "/api/v1/user/role/list/:uid"},
		},
	})

	list = append(list, &role.Api{
		Id:    10200,
		Label: "角色",
		Children: []*role.Api{
			{Id: 10201, Method: MethodGet, Label: "角色列表", Path: "/api/v1/role/list"},
			{Id: 10202, Method: MethodGet, Label: "角色信息", Path: "/api/v1/role/:id"},
			{Id: 10203, Method: MethodPos, Label: "添加角色", Path: "/api/v1/role"},
			{Id: 10204, Method: MethodPut, Label: "编辑角色", Path: "/api/v1/role/:id"},
			{Id: 10205, Method: MethodDel, Label: "删除角色", Path: "/api/v1/role/:id"},
		},
	})

	list = append(list, &role.Api{
		Id:    10300,
		Label: "角色权限",
		Children: []*role.Api{
			{Id: 10301, Method: MethodGet, Label: "规则列表", Path: "/api/v1/role/rule/list"},
			{Id: 10302, Method: MethodPos, Label: "添加角色规则", Path: "/api/v1/role/rules"},
			{Id: 10303, Method: MethodGet, Label: "角色规则列表", Path: "/api/v1/role/rules/:role_id"},
			{Id: 10311, Method: MethodGet, Label: "菜单列表", Path: "/api/v1/role/menu/list"},
			{Id: 10312, Method: MethodPos, Label: "添加角色菜单", Path: "/api/v1/role/menus"},
			{Id: 10313, Method: MethodGet, Label: "角色菜单列表", Path: "/api/v1/role/menus/:role_id"},
		},
	})

	list = append(list, &role.Api{
		Id:    90000,
		Label: "定时任务管理",
		Children: []*role.Api{
			{Id: 90001, Method: MethodGet, Label: "定时任务列表", Path: "/api/v1/cron/list"},
			{Id: 90002, Method: MethodPos, Label: "添加定时任务", Path: "/api/v1/cron"},
			{Id: 90003, Method: MethodPut, Label: "编辑定时任务", Path: "/api/v1/cron/:id"},
			{Id: 90004, Method: MethodDel, Label: "删除定时任务", Path: "/api/v1/cron/:id"},
		},
	})

	ret := &role.RuleListRet{
		List: list,
	}

	*rsp = *ret
	return nil
}
