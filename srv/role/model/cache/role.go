package cache

import (
	proto "github.com/isfk/aio/proto/role"
	"github.com/isfk/aio/srv/role/model/db"
)

// Cache struct
type Cache struct{}

// Model Model
var Model = &Cache{}

// Create Create
func (c *Cache) Create(info *proto.Role) (role *proto.Role, err error) {
	role, err = db.Model.Create(info)

	if err != nil {
		return nil, err
	}

	return role, nil
}

// Update Update
func (c *Cache) Update(info *proto.Role) (role *proto.Role, err error) {
	role, err = db.Model.Update(info)

	if err != nil {
		return nil, err
	}

	return role, nil
}

// Delete Delete
func (c *Cache) Delete(info *proto.Role) (role *proto.Role, err error) {
	role, err = db.Model.Delete(info)

	if err != nil {
		return nil, err
	}

	return role, nil
}

// InfoByID InfoByID
func (c *Cache) InfoByID(ID *proto.ID) (role *proto.Role, err error) {
	role, err = db.Model.InfoByID(ID)

	if err != nil {
		return nil, err
	}

	return role, nil
}

// InfoByName InfoByName
func (c *Cache) InfoByName(Name *proto.Name) (role *proto.Role, err error) {
	role, err = db.Model.InfoByName(Name)

	if err != nil {
		return nil, err
	}

	return role, nil
}

// ListByPage ListByPage
func (c *Cache) ListByPage(in *proto.ListReq) (*proto.RoleList, error) {
	list, err := db.Model.ListByPage(in)

	if err != nil {
		return nil, err
	}

	return list, nil
}
