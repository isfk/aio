package user

import (
	"context"
	"net/http"

	"github.com/isfk/aio/srv/api/pkg/api"

	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/isfk/aio/srv/api/pkg/validate"
	proto "github.com/isfk/aio/srv/user/proto/user"
	"github.com/labstack/echo/v4"
)

// Register Register
func (v1 *v1) Register(c echo.Context) error {
	var args RegisterArgs
	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	xlog.Info("v1.Register", xlog.Any("args", args))

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 按用户名查询是否已经存在
	user, err := v1.u.InfoByUsername(context.Background(), &proto.Username{Username: args.Username})
	if err != nil {
		xlog.Error("v1.u.InfoByUsername", xlog.String("err", err.Error()))
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if user.Id > 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户名已存在"})
	}

	// 按邮箱查询是否已经存在
	user, err = v1.u.InfoByEmail(context.Background(), &proto.Email{Email: args.Email})
	if err != nil {
		xlog.Error("v1.u.InfoByEmail", xlog.String("err", err.Error()))
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if user.Id > 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "邮箱已存在"})
	}

	// 创建用户
	user, err = v1.u.Create(context.Background(), &proto.User{
		Username: args.Username,
		Email:    args.Email,
		Password: args.Password,
	})

	return c.JSON(http.StatusOK, &api.RetOk{Data: user, Message: "注册成功"})
}
