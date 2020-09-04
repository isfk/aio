package handler

import (
	"context"

	"github.com/isfk/aio/srv/user/model/cache"
	"github.com/isfk/aio/srv/user/proto/user"
	log "github.com/micro/go-micro/v2/logger"
)

// User User
type User struct{}

// Create Create
func (e *User) Create(ctx context.Context, req *user.User, rsp *user.User) error {
	log.Info("Received User.Create req: ", req)
	info, err := cache.Model.Create(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// Update Update
func (e *User) Update(ctx context.Context, req *user.User, rsp *user.User) error {
	log.Info("Received User.Update req: ", req)
	info, err := cache.Model.Update(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// Delete Delete
func (e *User) Delete(ctx context.Context, req *user.User, rsp *user.User) error {
	log.Info("Received User.Delete req: ", req)
	info, err := cache.Model.Delete(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// InfoByID InfoByID
func (e *User) InfoByID(ctx context.Context, req *user.ID, rsp *user.User) error {
	log.Info("Received User.InfoByID req: ", req)
	info, err := cache.Model.InfoByID(&user.ID{Id: req.Id})

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// InfoByUsername InfoByUsername
func (e *User) InfoByUsername(ctx context.Context, req *user.Username, rsp *user.User) error {
	log.Info("Received User.InfoByUsername req: ", req)
	info, err := cache.Model.InfoByUsername(&user.Username{Username: req.Username})

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// InfoByPhone InfoByPhone
func (e *User) InfoByPhone(ctx context.Context, req *user.Phone, rsp *user.User) error {
	log.Info("Received User.InfoByPhone req: ", req)
	info, err := cache.Model.InfoByPhone(&user.Phone{Phone: req.Phone})

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// InfoByEmail InfoByEmail
func (e *User) InfoByEmail(ctx context.Context, req *user.Email, rsp *user.User) error {
	log.Info("Received User.InfoByEmail req: ", req)
	info, err := cache.Model.InfoByEmail(&user.Email{Email: req.Email})

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}

// ListByPage ListByPage
func (e *User) ListByPage(ctx context.Context, req *user.ListReq, rsp *user.UserList) error {
	log.Info("Received User.ListByPage req: ", req)
	list, err := cache.Model.ListByPage(req)

	if err != nil {
		return err
	}

	*rsp = *list
	return nil
}

// UpdatePassword UpdatePassword
func (e *User) UpdatePassword(ctx context.Context, req *user.IDPassword, rsp *user.User) error {
	log.Info("Received User.UpdatePassword req: ", req)
	info, err := cache.Model.UpdatePassword(req)

	if err != nil {
		return err
	}

	*rsp = *info
	return nil
}
