package cache

import (
	"github.com/isfk/aio/srv/article/model/db"
	proto "github.com/isfk/aio/srv/article/proto/article"
)

// CreateCategory CreateCategory
func (c *Cache) CreateCategory(info *proto.ArticleCategory) (article *proto.ArticleCategory, err error) {
	article, err = db.Model.CreateCategory(info)

	if err != nil {
		return nil, err
	}

	return article, nil
}

// UpdateCategory UpdateCategory
func (c *Cache) UpdateCategory(info *proto.ArticleCategory) (article *proto.ArticleCategory, err error) {
	article, err = db.Model.UpdateCategory(info)

	if err != nil {
		return nil, err
	}

	return article, nil
}

// DeleteCategory DeleteCategory
func (c *Cache) DeleteCategory(info *proto.ArticleCategory) (article *proto.ArticleCategory, err error) {
	article, err = db.Model.DeleteCategory(info)

	if err != nil {
		return nil, err
	}

	return article, nil
}

// CategoryInfoByID CategoryInfoByID
func (c *Cache) CategoryInfoByID(ID *proto.ID) (category *proto.ArticleCategory, err error) {
	category, err = db.Model.CategoryInfoByID(ID)

	if err != nil {
		return nil, err
	}

	return category, nil
}

// CategoryListByPage CategoryListByPage
func (c *Cache) CategoryListByPage(in *proto.CategoryListReq) (*proto.ArticleCategoryList, error) {
	list, err := db.Model.CategoryListByPage(in)

	if err != nil {
		return nil, err
	}

	return list, nil
}
