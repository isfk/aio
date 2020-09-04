package user

import (
	"context"
	"net/http"

	"github.com/isfk/aio/srv/api/pkg/api"
	log "github.com/micro/go-micro/v2/logger"

	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	"github.com/isfk/aio/srv/api/pkg/validate"
	proto "github.com/isfk/aio/srv/user/proto/user"
	"github.com/labstack/echo/v4"
)

// Add Add
func (v1 *v1) Add(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args AddArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Infof("v1.Add args %v", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 按用户名查询是否已经存在
	user, err := v1.u.InfoByUsername(context.Background(), &proto.Username{Username: args.Username})
	if err != nil {
		log.Error("v1.u.InfoByUsername err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if user.Id > 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户名已存在"})
	}

	// 按手机号查询是否已经存在
	user, err = v1.u.InfoByPhone(context.Background(), &proto.Phone{Phone: args.Phone})
	if err != nil {
		log.Error("v1.u.InfoByPhone err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if user.Id > 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "手机号已存在"})
	}

	// 按邮箱查询是否已经存在
	user, err = v1.u.InfoByEmail(context.Background(), &proto.Email{Email: args.Email})
	if err != nil {
		log.Error("v1.u.InfoByEmail err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if user.Id > 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "邮箱已存在"})
	}

	info := &proto.User{
		Username: args.Username,
		Nickname: args.Nickname,
		Phone:    args.Phone,
		Email:    args.Email,
		Gender:   1,
		Status:   1,
	}

	switch args.Gender {
	case 1:
		info.Gender = proto.Gender_WOMAN
	case 2:
		info.Gender = proto.Gender_MAN
	default:
	case 3:
		info.Gender = proto.Gender_SECRET
	}

	switch args.Status {
	default:
	case 1:
		info.Status = proto.Status_NORMAL
	case -1:
		info.Status = proto.Status_DELETE
	}

	resp, err := v1.u.Create(context.Background(), info)

	if err != nil {
		log.Error("v1.u.Create err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "添加成功",
	})
}

// Edit Edit
func (v1 *v1) Edit(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args EditArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Infof("v1.Edit args %v", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 查询用户信息
	user, err := v1.u.InfoByID(context.Background(), &proto.ID{Id: args.ID})
	if err != nil {
		log.Error("v1.u.InfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if user.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户不存在"})
	}

	// 按用户名查询是否已经存在
	userOther, err := v1.u.InfoByUsername(context.Background(), &proto.Username{Username: args.Username})
	if err != nil {
		log.Error("v1.u.InfoByUsername err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if userOther.Id > 0 && user.Id != userOther.Id {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户名已存在"})
	}

	// 按手机号查询是否已经存在
	userOther, err = v1.u.InfoByPhone(context.Background(), &proto.Phone{Phone: args.Phone})
	if err != nil {
		log.Error("v1.u.InfoByPhone err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if userOther.Id > 0 && user.Id != userOther.Id {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "手机号已存在"})
	}

	// 按邮箱查询是否已经存在
	userOther, err = v1.u.InfoByEmail(context.Background(), &proto.Email{Email: args.Email})
	if err != nil {
		log.Error("v1.u.InfoByEmail err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if userOther.Id > 0 && user.Id != userOther.Id {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "邮箱已存在"})
	}

	if len(args.Username) > 0 {
		user.Username = args.Username
	}
	if len(args.Nickname) > 0 {
		user.Nickname = args.Nickname
	}
	if len(args.Phone) > 0 {
		user.Phone = args.Phone
	}
	if len(args.Email) > 0 {
		user.Email = args.Email
	}

	switch args.Gender {
	case 1:
		user.Gender = proto.Gender_WOMAN
	case 2:
		user.Gender = proto.Gender_MAN
	case 3:
		user.Gender = proto.Gender_SECRET
	}

	switch args.Status {
	case 1:
		user.Status = proto.Status_NORMAL
	case -1:
		user.Status = proto.Status_DELETE
	}

	resp, err := v1.u.Update(context.Background(), user)

	if err != nil {
		log.Error("v1.u.Update err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "编辑成功",
	})
}

// Del Del
func (v1 *v1) Del(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args DelArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Infof("v1.Del args %v", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	user, err := v1.u.InfoByID(context.Background(), &proto.ID{Id: args.ID})

	if err != nil {
		log.Error("v1.u.InfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if user.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户不存在"})
	}

	resp, err := v1.u.Delete(context.Background(), user)

	if err != nil {
		log.Error("v1.u.Delete err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "删除成功",
	})
}

// Info Info
func (v1 *v1) Info(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args InfoArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}

	if args.ID == 0 {
		args.ID = userInfo.Id
	}

	log.Infof("v1.Info args %v", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	resp, err := v1.u.InfoByID(context.Background(), &proto.ID{
		Id: args.ID,
	})

	if err != nil {
		log.Error("v1.u.InfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "获取成功",
	})
}

// List List
func (v1 *v1) List(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args ListArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Infof("v1.List args %v", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	resp, err := v1.u.ListByPage(context.Background(), &proto.ListReq{
		Page:   args.Page,
		Limit:  args.Limit,
		Status: args.Status,
		Name:   args.Name,
	})

	if err != nil {
		log.Error("v1.u.ListByPage err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "获取成功",
	})
}
