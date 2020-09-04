package handler

import (
	"context"

	"github.com/isfk/aio/proto/article"
	"github.com/isfk/aio/srv/article/model/cache"
	log "github.com/micro/go-micro/v2/logger"
)

// Article Article
type Article struct{}

// Create Create
func (e *Article) Create(ctx context.Context, req *article.Article, rsp *article.Article) error {
	log.Info("Received Article.Create req: ", req)
	info, err := cache.Model.Create(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// Update Update
func (e *Article) Update(ctx context.Context, req *article.Article, rsp *article.Article) error {
	log.Info("Received Article.Update req: ", req)
	info, err := cache.Model.Update(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// Delete Delete
func (e *Article) Delete(ctx context.Context, req *article.Article, rsp *article.Article) error {
	log.Info("Received Article.Delete req: ", req)
	info, err := cache.Model.Delete(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// InfoByID InfoByID
func (e *Article) InfoByID(ctx context.Context, req *article.ID, rsp *article.Article) error {
	log.Info("Received Article.InfoByID req: ", req)
	info, err := cache.Model.InfoByID(&article.ID{Id: req.Id})

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// ListByPage ListByPage
func (e *Article) ListByPage(ctx context.Context, req *article.ListReq, rsp *article.ArticleList) error {
	log.Info("Received Article.ListByPage req: ", req)
	list, err := cache.Model.ListByPage(req)

	if err != nil {
		return err
	}

	*rsp = *list
	return nil
}
