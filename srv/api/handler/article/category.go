package article

import (
	"context"
	"net/http"

	"github.com/isfk/aio/srv/api/pkg/api"
	"github.com/isfk/aio/srv/api/pkg/validate"
	"github.com/labstack/echo/v4"

	proto "github.com/isfk/aio/proto/article"
	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	log "github.com/micro/go-micro/v2/logger"
)

// AddCategory AddCategory
func (v1 *v1) AddCategory(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args AddCategoryArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}

	log.Info("v1.article.AddCategory args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	info := &proto.ArticleCategory{
		UserId: userInfo.Id,
		Name:   args.Name,
	}

	resp, err := v1.a.CreateCategory(context.Background(), info)
	if err != nil {
		log.Error("v1.a.CreateCategory err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "添加成功",
	})
}

// EditCategory EditCategory
func (v1 *v1) EditCategory(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args EditCategoryArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}

	log.Info("v1.article.EditCategory args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 查询信息
	category, err := v1.a.CategoryInfoByID(context.Background(), &proto.ID{Id: args.ID})
	if err != nil {
		log.Error("v1.a.InfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if category.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "分类不存在"})
	}

	category.Name = args.Name

	resp, err := v1.a.UpdateCategory(context.Background(), category)
	if err != nil {
		log.Error("v1.a.UpdateCategory err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "编辑成功",
	})
}

// Del Del
func (v1 *v1) DelCategory(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args DelCategoryArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.article.Del args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	category, err := v1.a.CategoryInfoByID(context.Background(), &proto.ID{Id: args.ID})

	if err != nil {
		log.Error("v1.a.CategoryInfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	resp, err := v1.a.DeleteCategory(context.Background(), category)

	if err != nil {
		log.Error("v1.a.DeleteCategory err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "删除成功",
	})
}

// CategoryInfo CategoryInfo
func (v1 *v1) CategoryInfo(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args CategoryInfoArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.article.CategoryInfo args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	resp, err := v1.a.CategoryInfoByID(context.Background(), &proto.ID{
		Id: args.ID,
	})

	if err != nil {
		log.Error("v1.a.CategoryInfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "获取成功",
	})
}

// CategoryList CategoryList
func (v1 *v1) CategoryList(c echo.Context) error {
	userInfo := userMiddleware.GetUserInfo(c)
	if userInfo.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "未登录"})
	}

	var args CategoryListArgs

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err.Error()})
	}
	log.Info("v1.article.CategoryList args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	resp, err := v1.a.CategoryListByPage(context.Background(), &proto.CategoryListReq{
		Page:   args.Page,
		Limit:  args.Limit,
		Status: args.Status,
		Name:   args.Name,
		UserId: args.UserID,
	})

	if err != nil {
		log.Error("v1.a.CategoryListByPage err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "获取成功",
	})
}
