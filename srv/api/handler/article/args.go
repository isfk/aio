package article

// AddArgs struct
type AddArgs struct {
	CategoryID int64  `json:"category_id" form:"category_id" query:"category_id" validate:"required,number"`
	Title      string `json:"title" form:"title" query:"title" validate:"required,min=3,max=60"`
	Content    string `json:"content" form:"content" query:"content" validate:"required,min=3"`
	Status     int64  ` json:"status" form:"status" query:"status" validate:"required,oneof=1 -1"`
}

// EditArgs struct
type EditArgs struct {
	ID         int64  `json:"id" form:"id" query:"id" validate:"required,number"`
	CategoryID int64  `json:"category_id" form:"category_id" query:"category_id" validate:"required,number"`
	Title      string `json:"title" form:"title" query:"title" validate:"required,min=3,max=60"`
	Content    string `json:"content" form:"content" query:"content" validate:"required,min=3"`
	Status     int64  ` json:"status" form:"status" query:"status" validate:"required,oneof=1 -1"`
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
	Page       int64  `json:"page" form:"page" query:"page" validate:""`
	Limit      int64  `json:"limit" form:"limit" query:"limit" validate:""`
	Status     int64  `json:"status" form:"status" query:"status" validate:""`
	Title      string `json:"title" form:"title" query:"title" validate:""`
	UserID     int64  `json:"user_id" form:"user_id" query:"user_id" validate:""`
	CategoryID int64  `json:"category_id" form:"category_id" query:"category_id" validate:""`
}

// AddCategoryArgs struct
type AddCategoryArgs struct {
	Name string `json:"name" form:"name" query:"name" validate:"required,min=2,max=60"`
}

// EditCategoryArgs struct
type EditCategoryArgs struct {
	ID   int64  `json:"id" form:"id" query:"id" validate:"required,number"`
	Name string `json:"name" form:"name" query:"name" validate:"required,min=2,max=60"`
}

// DelCategoryArgs struct
type DelCategoryArgs struct {
	ID int64 `json:"id" form:"id" query:"id" validate:"required,number"`
}

// CategoryInfoArgs struct
type CategoryInfoArgs struct {
	ID int64 `json:"id" form:"id" query:"id" validate:""`
}

// CategoryListArgs struct
type CategoryListArgs struct {
	Page       int64  `json:"page" form:"page" query:"page" validate:""`
	Limit      int64  `json:"limit" form:"limit" query:"limit" validate:""`
	Status     int64  `json:"status" form:"status" query:"status" validate:""`
	Name       string `json:"name" form:"name" query:"name" validate:""`
	UserID     int64  `json:"user_id" form:"user_id" query:"user_id" validate:""`
	CategoryID int64  `json:"category_id" form:"category_id" query:"category_id" validate:""`
}

// EmptyArgs EmptyArgs
type EmptyArgs struct{}
