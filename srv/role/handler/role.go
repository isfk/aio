package handler

import (
	"context"

	"github.com/isfk/aio/proto/role"
	"github.com/isfk/aio/srv/role/model/cache"
	log "github.com/micro/go-micro/v2/logger"
)

// Role struct
type Role struct{}

// Create Create
func (e *Role) Create(ctx context.Context, req *role.Role, rsp *role.Role) error {
	log.Info("Received Role.Create req: ", req)
	info, err := cache.Model.Create(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// Update Update
func (e *Role) Update(ctx context.Context, req *role.Role, rsp *role.Role) error {
	log.Info("Received Role.Update req: ", req)
	info, err := cache.Model.Update(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// Delete Delete
func (e *Role) Delete(ctx context.Context, req *role.Role, rsp *role.Role) error {
	log.Info("Received Role.Delete req: ", req)
	info, err := cache.Model.Delete(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// InfoByID InfoByID
func (e *Role) InfoByID(ctx context.Context, req *role.ID, rsp *role.Role) error {
	log.Info("Received Role.InfoByID req: ", req)
	info, err := cache.Model.InfoByID(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// InfoByName InfoByName
func (e *Role) InfoByName(ctx context.Context, req *role.Name, rsp *role.Role) error {
	log.Info("Received Role.InfoByName req: ", req)
	info, err := cache.Model.InfoByName(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// ListByPage ListByPage
func (e *Role) ListByPage(ctx context.Context, req *role.ListReq, rsp *role.RoleList) error {
	log.Info("Received Role.ListByPage req: ", req)
	list, err := cache.Model.ListByPage(req)

	if err != nil {
		return err
	}

	*rsp = *list
	return nil
}
