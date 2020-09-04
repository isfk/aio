package user

// TokenRet struct
type TokenRet struct {
	Token string `json:"token" form:"token" query:"token"`
}

// LoginArgs struct
type LoginArgs struct {
	Username string `json:"username" form:"username" query:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" form:"password" query:"password" validate:"required,min=6,max=20"`
}

// RegisterArgs struct
type RegisterArgs struct {
	Username string `json:"username" form:"username" query:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" form:"email" query:"email" validate:"required,min=6,max=50,email"`
	Password string `json:"password" form:"password" query:"password" validate:"required,min=6,max=20"`
}

// AddArgs struct
type AddArgs struct {
	Username string `json:"username" form:"username" query:"username" validate:"required,min=3,max=20"`
	Nickname string `json:"nickname" form:"nickname" query:"nickname" validate:"required,min=3,max=20"`
	Phone    string `json:"phone" form:"phone" query:"phone" validate:"required,phone,len=11"`
	Email    string `json:"email" form:"email" query:"email" validate:"required,email"`
	Gender   int64  `json:"gender" form:"gender" query:"gender" validate:"required,oneof=1 2 3"`
	Status   int64  ` json:"status" form:"status" query:"status" validate:"required,oneof=1 -1"`
}

// EditArgs struct
type EditArgs struct {
	ID       int64  `json:"id" form:"id" query:"id" validate:"required,number"`
	Username string `json:"username" form:"username" query:"username" validate:"required,min=3,max=20"`
	Nickname string `json:"nickname" form:"nickname" query:"nickname" validate:"required,min=3,max=20"`
	Phone    string `json:"phone" form:"phone" query:"phone" validate:"required,phone,len=11"`
	Email    string `json:"email" form:"email" query:"email" validate:"required,email"`
	Gender   int64  `json:"gender" form:"gender" query:"gender" validate:"required,oneof=1 2 3"`
	Status   int64  ` json:"status" form:"status" query:"status" validate:"required,oneof=1 -1"`
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
	Page   int64  `json:"page" form:"page" query:"page" validate:""`
	Limit  int64  `json:"limit" form:"limit" query:"limit" validate:""`
	Status int64  `json:"status" form:"status" query:"status" validate:""`
	Name   string `json:"name" form:"name" query:"name" validate:""`
}

// AddRoleArgs struct
type AddRoleArgs struct {
	UserID int64 `json:"user_id" form:"user_id" query:"user_id" validate:"required,number"`
	RoleID int64 `json:"role_id" form:"role_id" query:"role_id" validate:"required,number"`
}

// DelRoleArgs struct
type DelRoleArgs struct {
	UserID int64 `json:"user_id" form:"user_id" query:"user_id" validate:"required,number"`
	RoleID int64 `json:"role_id" form:"role_id" query:"role_id" validate:"required,number"`
}

// RoleListArgs struct
type RoleListArgs struct {
	UserID int64 `json:"user_id" form:"user_id" query:"user_id" validate:"required,number"`
}

// EmptyArgs EmptyArgs
type EmptyArgs struct{}
