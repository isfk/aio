package db

import (
	"time"

	"github.com/isfk/aio/pkg/model"
	proto "github.com/isfk/aio/srv/role/proto/role"
)

// Mysql struct
type Mysql struct{}

// Model Model
var Model = &Mysql{}

// Create Create
func (m *Mysql) Create(info *proto.Role) (*proto.Role, error) {
	info.CreatedAt = time.Now().Unix()
	info.UpdatedAt = info.CreatedAt

	err := model.UseDB().Create(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Update Update
func (m *Mysql) Update(info *proto.Role) (*proto.Role, error) {
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Save(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Delete Delete
func (m *Mysql) Delete(info *proto.Role) (*proto.Role, error) {
	info.Status = proto.Status_DELETE
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Save(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// InfoByID InfoByID
func (m *Mysql) InfoByID(ID *proto.ID) (role *proto.Role, err error) {
	role = &proto.Role{}
	model.UseDB().Where("id = ?", ID.Id).First(&role)
	return role, nil
}

// InfoByName InfoByName
func (m *Mysql) InfoByName(Name *proto.Name) (role *proto.Role, err error) {
	role = &proto.Role{}
	model.UseDB().Where("name = ?", Name.Name).First(&role)
	return role, nil
}

// ListByPage ListByPage
func (m *Mysql) ListByPage(in *proto.ListReq) (r *proto.RoleList, err error) {
	r = &proto.RoleList{}
	db := model.UseDB()

	if in.Limit == 0 {
		in.Limit = 10
	}

	offset := (in.Page - 1) * in.Limit

	if in.Status == 1 || in.Status == -1 {
		db = db.Where("status = ?", in.Status)
	}

	var list []*proto.Role
	db.Offset(int(offset)).Limit(int(in.Limit)).Order("id DESC").Find(&list)

	var count int64
	db.Model(&proto.Role{}).Count(&count)

	r.List = list
	r.Count = count

	return r, nil
}
