package cache

import (
	proto "github.com/isfk/aio/proto/user"
	"github.com/isfk/aio/srv/user/model/db"
)

// Cache struct
type Cache struct{}

// Model Model
var Model = &Cache{}

// Create Create
func (c *Cache) Create(info *proto.User) (user *proto.User, err error) {
	user, err = db.Model.Create(info)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update Update
func (c *Cache) Update(info *proto.User) (user *proto.User, err error) {
	user, err = db.Model.Update(info)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete Delete
func (c *Cache) Delete(info *proto.User) (user *proto.User, err error) {
	user, err = db.Model.Delete(info)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// InfoByID InfoByID
func (c *Cache) InfoByID(ID *proto.ID) (user *proto.User, err error) {
	user, err = db.Model.InfoByID(ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// InfoByUsername InfoByUsername
func (c *Cache) InfoByUsername(Username *proto.Username) (user *proto.User, err error) {
	user, err = db.Model.InfoByUsername(Username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// InfoByPhone InfoByPhone
func (c *Cache) InfoByPhone(Phone *proto.Phone) (user *proto.User, err error) {
	user, err = db.Model.InfoByPhone(Phone)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// InfoByEmail InfoByEmail
func (c *Cache) InfoByEmail(Email *proto.Email) (user *proto.User, err error) {
	user, err = db.Model.InfoByEmail(Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// ListByPage ListByPage
func (c *Cache) ListByPage(in *proto.ListReq) (*proto.UserList, error) {
	list, err := db.Model.ListByPage(in)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// UpdatePassword UpdatePassword
func (c *Cache) UpdatePassword(info *proto.IDPassword) (user *proto.User, err error) {
	user, err = db.Model.UpdatePassword(info)

	if err != nil {
		return nil, err
	}

	return user, nil
}
