package db

import (
	"time"

	"github.com/isfk/aio/pkg/model"
	proto "github.com/isfk/aio/proto/article"
)

// CreateCategory CreateCategory
func (m *DB) CreateCategory(info *proto.ArticleCategory) (*proto.ArticleCategory, error) {
	info.CreatedAt = time.Now().Unix()
	info.UpdatedAt = info.CreatedAt

	err := model.UseDB().Create(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// UpdateCategory UpdateCategory
func (m *DB) UpdateCategory(info *proto.ArticleCategory) (*proto.ArticleCategory, error) {
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Save(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// DeleteCategory DeleteCategory
func (m *DB) DeleteCategory(info *proto.ArticleCategory) (*proto.ArticleCategory, error) {
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Where("id = ?", info.Id).Delete(&proto.ArticleCategory{}).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// CategoryInfoByID CategoryInfoByID
func (m *DB) CategoryInfoByID(ID *proto.ID) (info *proto.ArticleCategory, err error) {
	info = &proto.ArticleCategory{}
	model.UseDB().Where("id = ?", ID.Id).First(&info)
	return info, nil
}

// CategoryListByPage CategoryListByPage
func (m *DB) CategoryListByPage(in *proto.CategoryListReq) (r *proto.ArticleCategoryList, err error) {
	r = &proto.ArticleCategoryList{}
	db := model.UseDB()

	if in.Limit == 0 {
		in.Limit = 10
	}

	offset := (in.Page - 1) * in.Limit

	if len(in.Name) > 0 {
		db = db.Where("name like ?", "%"+in.Name+"%")
	}

	if in.Uid > 0 {
		db = db.Where("uid = ?", in.Uid)
	}

	var list []*proto.ArticleCategory
	db.Offset(int(offset)).Limit(int(in.Limit)).Order("id DESC").Find(&list)

	var count int64
	db.Model(&proto.ArticleCategory{}).Count(&count)

	r.List = list
	r.Count = count

	return r, nil
}
