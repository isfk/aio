package user

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	proto "github.com/isfk/aio/srv/user/proto/user"
	"github.com/labstack/echo/v4"
)

const (
	// InfoContextKey InfoContextKey
	InfoContextKey = "ypj_user_info"
)

// NeedLogin 需要登录
func NeedLogin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			if claims["user_info"] != nil {
				userInfo := claims["user_info"].(map[string]interface{})
				c.Set(InfoContextKey, userInfo)
				return next(c)
			}

			return echo.NewHTTPError(http.StatusBadRequest, "未登录")
		}
	}
}

// GetUserInfo 获取当前用户信息
func GetUserInfo(c echo.Context) (userInfo *proto.User) {
	userInfoFromContext := c.Get(InfoContextKey)
	userJSON, err := json.Marshal(userInfoFromContext.(map[string]interface{}))
	if err != nil {
		return &proto.User{}
	}
	err = json.Unmarshal([]byte(userJSON), &userInfo)
	if err != nil {
		return &proto.User{}
	}
	return userInfo
}
