package handler

import (
	"context"

	"github.com/isfk/aio/proto/user"
	"github.com/isfk/aio/srv/user/model/cache"
	log "github.com/micro/go-micro/v2/logger"
)

// CreateUserRole CreateUserRole
func (e *User) CreateUserRole(ctx context.Context, req *user.UserRole, rsp *user.UserRole) error {
	log.Info("Received User.CreateUserRole req: ", req)
	info, err := cache.Model.CreateUserRole(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// DeleteUserRole DeleteUserRole
func (e *User) DeleteUserRole(ctx context.Context, req *user.UserRole, rsp *user.UserRole) error {
	log.Info("Received User.DeleteUserRole req: ", req)
	info, err := cache.Model.DeleteUserRole(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// UserRoleInfo UserRoleInfo
func (e *User) UserRoleInfo(ctx context.Context, req *user.UserRole, rsp *user.UserRole) error {
	log.Info("Received User.UserRoleInfo req: ", req)
	info, err := cache.Model.UserRoleInfo(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// UserRoleListByUID UserRoleListByUID
func (e *User) UserRoleListByUID(ctx context.Context, req *user.UID, rsp *user.UserRoleList) error {
	log.Info("Received User.UserRoleListByUID req: ", req)
	list, err := cache.Model.UserRoleListByUID(req)

	if err != nil {
		return err
	}

	*rsp = *list
	return nil
}
