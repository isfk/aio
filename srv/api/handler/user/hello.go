package user

import (
	"net/http"

	"github.com/isfk/aio/srv/api/pkg/api"
	"github.com/isfk/aio/srv/api/pkg/validate"
	"github.com/labstack/echo/v4"
	log "github.com/micro/go-micro/v2/logger"
)

// Args struct
type Args struct {
	Name string `json:"name" form:"name" query:"name" validate:"required"`
}

// Logout Logout
func (v1 *v1) SayHello(c echo.Context) error {
	var args Args
	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.SayHello args", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	return c.JSON(http.StatusOK, &api.RetOk{Message: "hello, " + args.Name})
}
