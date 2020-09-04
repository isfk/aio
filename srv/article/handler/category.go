package handler

import (
	"context"

	"github.com/isfk/aio/proto/article"
	"github.com/isfk/aio/srv/article/model/cache"
	log "github.com/micro/go-micro/v2/logger"
)

// CreateCategory CreateCategory
func (e *Article) CreateCategory(ctx context.Context, req *article.ArticleCategory, rsp *article.ArticleCategory) error {
	log.Info("Received Article.CreateCategory req: ", req)
	info, err := cache.Model.CreateCategory(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// UpdateCategory UpdateCategory
func (e *Article) UpdateCategory(ctx context.Context, req *article.ArticleCategory, rsp *article.ArticleCategory) error {
	log.Info("Received Article.UpdateCategory req: ", req)
	info, err := cache.Model.UpdateCategory(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// DeleteCategory DeleteCategory
func (e *Article) DeleteCategory(ctx context.Context, req *article.ArticleCategory, rsp *article.ArticleCategory) error {
	log.Info("Received Article.DeleteCategory req: ", req)
	info, err := cache.Model.DeleteCategory(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// CategoryInfoByID CategoryInfoByID
func (e *Article) CategoryInfoByID(ctx context.Context, req *article.ID, rsp *article.ArticleCategory) error {
	log.Info("Received Article.CategoryInfoByID req: ", req)
	info, err := cache.Model.CategoryInfoByID(&article.ID{Id: req.Id})

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// CategoryListByPage CategoryListByPage
func (e *Article) CategoryListByPage(ctx context.Context, req *article.CategoryListReq, rsp *article.ArticleCategoryList) error {
	log.Info("Received Article.CategoryListByPage req: ", req)
	list, err := cache.Model.CategoryListByPage(req)

	if err != nil {
		return err
	}

	*rsp = *list
	return nil
}
