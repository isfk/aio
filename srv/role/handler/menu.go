package handler

import (
	"context"

	"github.com/isfk/aio/proto/role"
)

// MenuList MenuList
func (e *Role) MenuList(ctx context.Context, req *role.Null, rsp *role.MenuListRet) error {
	var list []*role.Menu
	list = append(list, &role.Menu{
		Id:        10100,
		Label:     "用户管理",
		Path:      "/user",
		Component: "layout",
		Hidden:    false,
		Redirect:  "/user/list",
		Meta:      &role.Meta{Title: "用户管理", Icon: "el-icon-user-solid"},
		Children: []*role.Menu{
			{
				Id: 10101, Label: "用户列表", Path: "list", Name: "user_list", Component: "user_list", Meta: &role.Meta{Title: "用户列表"},
			},
		},
	})

	list = append(list, &role.Menu{
		Id:        10200,
		Label:     "角色管理",
		Path:      "/role",
		Component: "layout",
		Hidden:    false,
		Redirect:  "/role/list",
		Meta:      &role.Meta{Title: "角色管理", Icon: "el-icon-user-solid"},
		Children: []*role.Menu{
			{
				Id: 10201, Label: "角色列表", Path: "list", Name: "role_list", Component: "role_list", Meta: &role.Meta{Title: "角色列表"},
			},
		},
	})

	list = append(list, &role.Menu{
		Id:        10400,
		Label:     "文章管理",
		Path:      "/article",
		Component: "layout",
		Hidden:    false,
		Redirect:  "/article/list",
		Meta:      &role.Meta{Title: "文章管理", Icon: "el-icon-user-solid"},
		Children: []*role.Menu{
			{
				Id: 10401, Label: "文章列表", Path: "list", Name: "article_list", Component: "article_list", Meta: &role.Meta{Title: "文章列表"}, Children: []*role.Menu{},
			},
			{
				Id: 10451, Label: "分类列表", Path: "category/list", Name: "category_list", Component: "category_list", Meta: &role.Meta{Title: "分类列表"}, Children: []*role.Menu{},
			},
		},
	})

	ret := &role.MenuListRet{
		List: list,
	}

	*rsp = *ret
	return nil
}
