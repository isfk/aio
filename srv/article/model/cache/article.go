package cache

import (
	"github.com/isfk/aio/srv/article/model/db"
	proto "github.com/isfk/aio/srv/article/proto/article"
)

// Cache struct
type Cache struct{}

// Model Model
var Model = &Cache{}

// Create Create
func (c *Cache) Create(info *proto.Article) (article *proto.Article, err error) {
	article, err = db.Model.Create(info)

	if err != nil {
		return nil, err
	}

	return article, nil
}

// Update Update
func (c *Cache) Update(info *proto.Article) (article *proto.Article, err error) {
	article, err = db.Model.Update(info)

	if err != nil {
		return nil, err
	}

	return article, nil
}

// Delete Delete
func (c *Cache) Delete(info *proto.Article) (article *proto.Article, err error) {
	article, err = db.Model.Delete(info)

	if err != nil {
		return nil, err
	}

	return article, nil
}

// InfoByID InfoByID
func (c *Cache) InfoByID(ID *proto.ID) (article *proto.Article, err error) {
	article, err = db.Model.InfoByID(ID)

	if err != nil {
		return nil, err
	}

	return article, nil
}

// ListByPage ListByPage
func (c *Cache) ListByPage(in *proto.ListReq) (*proto.ArticleList, error) {
	list, err := db.Model.ListByPage(in)

	if err != nil {
		return nil, err
	}

	return list, nil
}
