package api

// RetErr 接口错误返回
type RetErr struct {
	Code 	int64  	`json:"code" form:"code" query:"code"`
	Message string	`json:"message" form:"message" query:"message"`
}

// RetOk 接口正常返回
type RetOk struct {
	Data 	interface{}	`json:"data" form:"data" query:"data"`
	Message string		`json:"message" form:"message" query:"message"`
}
