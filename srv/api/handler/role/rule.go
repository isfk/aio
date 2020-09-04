package role

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	casbinV2 "github.com/casbin/casbin/v2"
	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	"github.com/isfk/aio/srv/api/pkg/api"
	rolePkg "github.com/isfk/aio/srv/api/pkg/role"
	"github.com/isfk/aio/srv/api/pkg/validate"
	"github.com/isfk/aio/srv/role/proto/role"
	"github.com/labstack/echo/v4"
	log "github.com/micro/go-micro/v2/logger"
)

// RuleList RuleList
func (v1 *v1) RuleList(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	ruleList, err := v1.r.RuleList(context.Background(), &role.Null{})
	if err != nil {
		log.Error("v1.r.RuleList err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    ruleList.List,
		Message: "获取成功",
	})
}

// AddRoleRule 添加角色规则
func (v1 *v1) AddRoleRule(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args AddRoleRuleArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.casbin.AddRoleRule args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	roleIDString := strconv.FormatInt(args.RoleID, 10)

	v1.e.RemoveFilteredPolicy(0, fmt.Sprintf("role_id:%s.api", roleIDString))

	ruleList, err := v1.r.RuleList(context.Background(), &role.Null{})
	if err != nil {
		log.Error("v1.r.InfoByUsername err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	apiListAll := rolePkg.GetRuleListAll(ruleList.List)

	go func(e *casbinV2.Enforcer, IDList []int64) {
		for _, v := range args.IDList {
			if a, ok := apiListAll[v]; ok {
				ok, err := v1.e.AddPolicy(fmt.Sprintf("role_id:%s.api", roleIDString), a.Path, a.Method, strconv.FormatInt(v, 10))

				if err != nil {
					log.Error("v1.e.AddPolicy err ", err.Error())
				}

				if ok == false {
					log.Error("v1.e.AddPolicy err ", "添加失败")
				}
			}
		}
	}(v1.e, args.IDList)

	return c.JSON(http.StatusOK, &api.RetOk{
		Message: "添加成功",
	})
}

// RoleRules RoleRules
func (v1 *v1) RoleRules(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args RulesArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.AddRole args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    map[string]interface{}{"list": v1.e.GetFilteredPolicy(0, fmt.Sprintf("role_id:%d.api", args.RoleID))},
		Message: "获取成功",
	})
}
