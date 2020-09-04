package tweet

import (
	"context"
	"fmt"
	"net/http"

	"github.com/isfk/aio/srv/api/pkg/api"
	"github.com/isfk/aio/srv/api/pkg/validate"
	"github.com/labstack/echo/v4"

	tweet "github.com/isfk/aio/proto/tweet"
)

// Args struct
type Args struct {
	Name string `json:"name" form:"name" query:"name"`
}

// TweetCall TweetCall
func (v1 *v1) TweetCall(c echo.Context) error {
	var args Args

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}

	fmt.Println("args", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	rsp, err := v1.Client.Call(context.TODO(), &tweet.Request{
		Name: args.Name,
	})

	if err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	c.JSON(http.StatusOK, rsp)
	return nil
}
