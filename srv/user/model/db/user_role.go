package db

import (
	"time"

	"github.com/isfk/aio/pkg/model"
	proto "github.com/isfk/aio/proto/user"
)

// CreateUserRole CreateUserRole
func (m *DB) CreateUserRole(info *proto.UserRole) (*proto.UserRole, error) {
	info.CreatedAt = time.Now().Unix()
	info.UpdatedAt = info.CreatedAt

	err := model.UseDB().Create(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// DeleteUserRole DeleteUserRole
func (m *DB) DeleteUserRole(info *proto.UserRole) (*proto.UserRole, error) {
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Where("user_id = ? AND role_id = ?", info.UserId, info.RoleId).Delete(&proto.UserRole{}).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// UserRoleInfo UserRoleInfo
func (m *DB) UserRoleInfo(info *proto.UserRole) (*proto.UserRole, error) {
	userRole := &proto.UserRole{}
	return userRole, nil
}

// UserRoleListByUserID UserRoleListByUserID
func (m *DB) UserRoleListByUserID(in *proto.UserID) (r *proto.UserRoleList, err error) {
	r = &proto.UserRoleList{}
	db := model.UseDB().Debug()

	if in.UserId > 0 {
		db = db.Where("uid = ?", in.UserId)
	}

	var list []*proto.UserRole
	db.Order("id DESC").Find(&list)

	r.List = list

	return r, nil
}
