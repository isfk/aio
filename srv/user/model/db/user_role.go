package db

import (
	"time"

	"github.com/isfk/aio/pkg/model"
	proto "github.com/isfk/aio/proto/user"
)

// CreateUserRole CreateUserRole
func (m *Mysql) CreateUserRole(info *proto.UserRole) (*proto.UserRole, error) {
	info.CreatedAt = time.Now().Unix()
	info.UpdatedAt = info.CreatedAt

	err := model.UseDB().Create(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// DeleteUserRole DeleteUserRole
func (m *Mysql) DeleteUserRole(info *proto.UserRole) (*proto.UserRole, error) {
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Where("uid = ? AND role_id = ?", info.Uid, info.RoleId).Delete(&proto.UserRole{}).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// UserRoleInfo UserRoleInfo
func (m *Mysql) UserRoleInfo(info *proto.UserRole) (*proto.UserRole, error) {
	userRole := &proto.UserRole{}
	return userRole, nil
}

// UserRoleListByUID UserRoleListByUID
func (m *Mysql) UserRoleListByUID(in *proto.UID) (r *proto.UserRoleList, err error) {
	r = &proto.UserRoleList{}
	db := model.UseDB().Debug()

	if in.Uid > 0 {
		db = db.Where("uid = ?", in.Uid)
	}

	var list []*proto.UserRole
	db.Order("id DESC").Find(&list)

	r.List = list

	return r, nil
}
