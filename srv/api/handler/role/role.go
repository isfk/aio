package role

import (
	"context"
	"net/http"

	"github.com/isfk/aio/srv/api/pkg/api"
	"github.com/isfk/aio/srv/api/pkg/validate"

	proto "github.com/isfk/aio/proto/role"
	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	"github.com/labstack/echo/v4"
	log "github.com/micro/go-micro/v2/logger"
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
	log.Info("v1.role.Add args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 按名称查询是否已经存在
	role, err := v1.r.InfoByName(context.Background(), &proto.Name{Name: args.Name})
	if err != nil {
		log.Error("v1.r.InfoByUsername err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if role.Id > 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "角色已存在"})
	}

	info := &proto.Role{
		Name: args.Name,
	}

	switch args.Status {
	case 1:
		info.Status = proto.Status_NORMAL
	case -1:
		info.Status = proto.Status_DELETE
	}

	resp, err := v1.r.Create(context.Background(), info)

	if err != nil {
		log.Error("v1.r.Create err ", err.Error())
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
	log.Info("v1.role.Edit args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 查询角色信息
	role, err := v1.r.InfoByID(context.Background(), &proto.ID{Id: args.ID})
	if err != nil {
		log.Error("v1.r.InfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if role.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "角色不存在"})
	}

	// 按用户名查询是否已经存在
	roleOther, err := v1.r.InfoByName(context.Background(), &proto.Name{Name: args.Name})
	if err != nil {
		log.Error("v1.r.InfoByName err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if roleOther.Id > 0 && role.Id != roleOther.Id {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户名已存在"})
	}

	role.Name = args.Name

	switch args.Status {
	case 1:
		role.Status = proto.Status_NORMAL
	case -1:
		role.Status = proto.Status_DELETE
	}

	resp, err := v1.r.Update(context.Background(), role)

	if err != nil {
		log.Error("v1.r.Update err ", err.Error())
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
	log.Info("v1.role.Del args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	role, err := v1.r.InfoByID(context.Background(), &proto.ID{Id: args.ID})

	if err != nil {
		log.Error("v1.r.InfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	resp, err := v1.r.Delete(context.Background(), role)

	if err != nil {
		log.Error("v1.r.Delete err ", err.Error())
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
	log.Info("v1.role.Info args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	resp, err := v1.r.InfoByID(context.Background(), &proto.ID{
		Id: args.ID,
	})

	if err != nil {
		log.Error("v1.r.InfoByID err ", err.Error())
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
	log.Info("v1.List args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	resp, err := v1.r.ListByPage(context.Background(), &proto.ListReq{
		Page:   args.Page,
		Limit:  args.Limit,
		Status: args.Status,
	})

	if err != nil {
		log.Error("v1.r.ListByPage err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "获取成功",
	})
}
