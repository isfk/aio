package cache

import (
	"github.com/isfk/aio/srv/user/model/db"
	proto "github.com/isfk/aio/srv/user/proto/user"
)

// CreateUserRole CreateUserRole
func (c *Cache) CreateUserRole(info *proto.UserRole) (userRole *proto.UserRole, err error) {
	userRole, err = db.Model.CreateUserRole(info)

	if err != nil {
		return nil, err
	}

	return userRole, nil
}

// DeleteUserRole DeleteUserRole
func (c *Cache) DeleteUserRole(info *proto.UserRole) (userRole *proto.UserRole, err error) {
	userRole, err = db.Model.DeleteUserRole(info)

	if err != nil {
		return nil, err
	}

	return userRole, nil
}

// UserRoleInfo UserRoleInfo
func (c *Cache) UserRoleInfo(info *proto.UserRole) (userRole *proto.UserRole, err error) {
	userRole, err = db.Model.UserRoleInfo(info)

	if err != nil {
		return nil, err
	}

	return userRole, nil
}

// UserRoleListByUID UserRoleListByUID
func (c *Cache) UserRoleListByUID(in *proto.UID) (*proto.UserRoleList, error) {
	list, err := db.Model.UserRoleListByUID(in)

	if err != nil {
		return nil, err
	}

	return list, nil
}
