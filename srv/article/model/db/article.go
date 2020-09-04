package db

import (
	"time"

	"github.com/isfk/aio/pkg/model"
	proto "github.com/isfk/aio/srv/article/proto/article"
)

// Mysql struct
type Mysql struct{}

// Model Model
var Model = &Mysql{}

// Create Create
func (m *Mysql) Create(info *proto.Article) (*proto.Article, error) {
	info.CreatedAt = time.Now().Unix()
	info.UpdatedAt = info.CreatedAt

	if info.Status == 0 {
		info.Status = 1
	}

	err := model.UseDB().Create(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Update Update
func (m *Mysql) Update(info *proto.Article) (*proto.Article, error) {
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Save(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Delete Delete
func (m *Mysql) Delete(info *proto.Article) (*proto.Article, error) {
	info.Status = proto.Status_DELETE
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Save(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// InfoByID InfoByID
func (m *Mysql) InfoByID(ID *proto.ID) (info *proto.Article, err error) {
	info = &proto.Article{}
	model.UseDB().Where("id = ?", ID.Id).First(&info)
	return info, nil
}

// ListByPage ListByPage
func (m *Mysql) ListByPage(in *proto.ListReq) (r *proto.ArticleList, err error) {
	r = &proto.ArticleList{}
	db := model.UseDB()

	if in.Limit == 0 {
		in.Limit = 10
	}

	offset := (in.Page - 1) * in.Limit

	if in.Status == 1 || in.Status == -1 {
		db = db.Where("status = ?", in.Status)
	}

	if len(in.Title) > 0 {
		db = db.Where("title like ?", "%"+in.Title+"%")
	}

	if in.Uid > 0 {
		db = db.Where("uid = ?", in.Uid)
	}

	if in.CategoryId > 0 {
		db = db.Where("category_id = ?", in.CategoryId)
	}

	var list []*proto.Article
	db.Offset(int(offset)).Limit(int(in.Limit)).Order("id DESC").Find(&list)

	var count int64
	db.Model(&proto.Article{}).Count(&count)

	r.List = list
	r.Count = count

	return r, nil
}
