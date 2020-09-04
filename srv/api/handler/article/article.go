package article

import (
	"context"
	"net/http"

	"github.com/isfk/aio/srv/api/pkg/api"
	"github.com/isfk/aio/srv/api/pkg/validate"
	"github.com/labstack/echo/v4"

	userMiddleware "github.com/isfk/aio/srv/api/middleware/user"
	proto "github.com/isfk/aio/srv/article/proto/article"
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

	log.Info("v1.article.Add args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 检查 category_id 是不是自己的
	categoryInfo, err := v1.a.CategoryInfoByID(context.Background(), &proto.ID{
		Id: args.CategoryID,
	})
	if err != nil {
		log.Error("v1.a.CategoryInfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if categoryInfo.Uid != userInfo.Id {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：分类错误"})
	}

	info := &proto.Article{
		UserId:     userInfo.Id,
		CategoryId: args.CategoryID,
		Title:      args.Title,
		Content:    args.Content,
	}

	switch args.Status {
	default:
	case 1:
		info.Status = proto.Status_NORMAL
	case -1:
		info.Status = proto.Status_DELETE
	}

	resp, err := v1.a.Create(context.Background(), info)
	if err != nil {
		log.Error("v1.a.Create err ", err.Error())
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

	log.Info("v1.article.Edit args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	// 检查 category_id 是不是自己的
	categoryInfo, err := v1.a.CategoryInfoByID(context.Background(), &proto.ID{
		Id: args.CategoryID,
	})
	if err != nil {
		log.Error("v1.a.CategoryInfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}
	if categoryInfo.Uid != userInfo.Id {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：分类错误"})
	}

	// 查询信息
	article, err := v1.a.InfoByID(context.Background(), &proto.ID{Id: args.ID})
	if err != nil {
		log.Error("v1.a.InfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	if article.Id == 0 {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "文章不存在"})
	}

	article.Title = args.Title
	article.Content = args.Content

	switch args.Status {
	case 1:
		article.Status = proto.Status_NORMAL
	case -1:
		article.Status = proto.Status_DELETE
	}

	resp, err := v1.a.Update(context.Background(), article)
	if err != nil {
		log.Error("v1.a.Update err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "添加成功",
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
	log.Info("v1.article.Del args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	article, err := v1.a.InfoByID(context.Background(), &proto.ID{Id: args.ID})

	if err != nil {
		log.Error("v1.a.InfoByID err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	resp, err := v1.a.Delete(context.Background(), article)

	if err != nil {
		log.Error("v1.a.Delete err ", err.Error())
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
	log.Info("v1.article.Info args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	resp, err := v1.a.InfoByID(context.Background(), &proto.ID{
		Id: args.ID,
	})

	if err != nil {
		log.Error("v1.a.InfoByID err ", err.Error())
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
	log.Info("v1.article.List args ", args)

	if err := validate.Validate(args); err != "" {
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: "参数错误：" + err})
	}

	resp, err := v1.a.ListByPage(context.Background(), &proto.ListReq{
		Page:       args.Page,
		Limit:      args.Limit,
		Status:     args.Status,
		Title:      args.Title,
		Uid:        args.UID,
		CategoryId: args.CategoryID,
	})

	if err != nil {
		log.Error("v1.a.ListByPage err ", err.Error())
		return c.JSON(http.StatusOK, &api.RetErr{Code: 40001, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &api.RetOk{
		Data:    resp,
		Message: "获取成功",
	})
}
