package user

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	proto "github.com/isfk/aio/proto/user"
	"github.com/isfk/aio/srv/api/pkg/api"
	"github.com/isfk/aio/srv/api/pkg/utils"
	"github.com/isfk/aio/srv/api/pkg/validate"
	"github.com/labstack/echo/v4"
	log "github.com/micro/go-micro/v2/logger"
)

// Login Login
func (v1 *v1) Login(c echo.Context) error {
	var args LoginArgs
	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.Login args", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 按用户名查询是否已经存在
	user, err := v1.u.InfoByUsername(context.Background(), &proto.Username{Username: args.Username})
	if err != nil {
		log.Info("v1.u.InfoByUsername err", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	log.Info("user", user)

	if user.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户名或密码错误"})
	}

	if user.Password != utils.HashPassword(args.Password, user.Salt) {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户名或密码错误"})
	}

	if user.Status != proto.Status_NORMAL {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "用户状态不可用"})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_info"] = user
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t := &TokenRet{}
	t.Token, err = token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &api.RetOk{Data: t, Message: "登录成功"})
}

// Logout Logout
func (v1 *v1) Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, &api.RetOk{Message: "退出成功"})
}
