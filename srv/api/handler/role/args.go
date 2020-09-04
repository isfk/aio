package role

// AddArgs struct
type AddArgs struct {
	Name   string `json:"name" form:"name" query:"name" validate:"required,min=2,max=20"`
	Status int64  `json:"status" form:"status" query:"status" validate:"required,oneof=1 -1"`
}

// EditArgs struct
type EditArgs struct {
	ID     int64  `json:"id" form:"id" query:"id" validate:"required,number"`
	Name   string `json:"name" form:"name" query:"name" validate:"required,min=2,max=20"`
	Status int64  `json:"status" form:"status" query:"status" validate:"required,oneof=1 -1"`
}

// DelArgs struct
type DelArgs struct {
	ID int64 `json:"id" form:"id" query:"id" validate:"required,number"`
}

// InfoArgs struct
type InfoArgs struct {
	ID int64 `json:"id" form:"id" query:"id" validate:""`
}

// ListArgs struct
type ListArgs struct {
	Page   int64 `json:"page" form:"page" query:"page" validate:""`
	Limit  int64 `json:"limit" form:"limit" query:"limit" validate:""`
	Status int64 `json:"status" form:"status" query:"status" validate:""`
}

// RulesArgs struct
type RulesArgs struct {
	RoleID int64 `json:"role_id" form:"role_id" query:"role_id" validate:"required,number"`
}

// AddRoleRuleArgs AddRoleRuleArgs
type AddRoleRuleArgs struct {
	RoleID int64   `json:"role_id" form:"role_id" query:"role_id" validate:"required,number"`
	IDList []int64 `json:"id_list" form:"id_list" query:"id_list" validate:"required"`
}

// MenusArgs struct
type MenusArgs struct {
	RoleID int64 `json:"role_id" form:"role_id" query:"role_id" validate:"required,number"`
}

// AddRoleMenuArgs AddRoleMenuArgs
type AddRoleMenuArgs struct {
	RoleID int64   `json:"role_id" form:"role_id" query:"role_id" validate:"required,number"`
	IDList []int64 `json:"id_list" form:"id_list" query:"id_list" validate:"required"`
}
