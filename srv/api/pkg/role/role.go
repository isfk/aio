package role

import "github.com/isfk/aio/proto/role"

// GetRuleListAll GetRuleListAll
func GetRuleListAll(apis []*role.Api) map[int64]*role.Api {
	listAll := map[int64]*role.Api{}
	for _, v := range apis {
		for _, c := range v.Children {
			listAll[c.Id] = c
		}
	}

	return listAll
}



// GetMenuListAll GetMenuListAll
func GetMenuListAll(menus []*role.Menu) map[int64]*role.Menu {
	listAll := map[int64]*role.Menu{}
	for _, v := range menus {
		listAll[v.Id] = &role.Menu{
			Id:   v.Id,
			Path: v.Path,
		}
		for _, c := range v.Children {
			listAll[c.Id] = c
		}
	}
	return listAll
}
