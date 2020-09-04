package db

import (
	"time"

	"github.com/isfk/aio/pkg/utils"

	"github.com/isfk/aio/pkg/model"
	proto "github.com/isfk/aio/srv/user/proto/user"
	"github.com/labstack/gommon/random"
)

// Mysql struct
type Mysql struct{}

// Model Model
var Model = &Mysql{}

// Create Create
func (m *Mysql) Create(info *proto.User) (*proto.User, error) {
	info.CreatedAt = time.Now().Unix()
	info.UpdatedAt = info.CreatedAt
	info.Salt = random.String(6)

	if len(info.Password) == 0 {
		info.Password = utils.HashPassword("123456", info.Salt)
	} else {
		info.Password = utils.HashPassword(info.Password, info.Salt)
	}
	if len(info.Nickname) == 0 {
		info.Nickname = info.Username
	}

	info.Status = 1

	err := model.UseDB().Create(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Update Update
func (m *Mysql) Update(info *proto.User) (*proto.User, error) {
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Omit("password", "salt").Save(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Delete Delete
func (m *Mysql) Delete(info *proto.User) (*proto.User, error) {
	info.Status = proto.Status_DELETE
	info.UpdatedAt = time.Now().Unix()
	err := model.UseDB().Save(info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// InfoByID InfoByID
func (m *Mysql) InfoByID(ID *proto.ID) (info *proto.User, err error) {
	info = &proto.User{}
	model.UseDB().Where("id = ?", ID.Id).First(&info)
	return info, nil
}

// InfoByUsername InfoByUsername
func (m *Mysql) InfoByUsername(Username *proto.Username) (info *proto.User, err error) {
	info = &proto.User{}
	model.UseDB().Where("username = ?", Username.Username).First(&info)
	return info, nil
}

// InfoByPhone InfoByPhone
func (m *Mysql) InfoByPhone(Phone *proto.Phone) (info *proto.User, err error) {
	info = &proto.User{}
	model.UseDB().Where("phone = ?", Phone.Phone).First(&info)
	return info, nil
}

// InfoByEmail InfoByEmail
func (m *Mysql) InfoByEmail(Email *proto.Email) (info *proto.User, err error) {
	info = &proto.User{}
	model.UseDB().Where("email = ?", Email.Email).First(&info)
	return info, nil
}

// ListByPage ListByPage
func (m *Mysql) ListByPage(in *proto.ListReq) (r *proto.UserList, err error) {
	r = &proto.UserList{}
	db := model.UseDB()

	if in.Limit == 0 {
		in.Limit = 10
	}

	offset := (in.Page - 1) * in.Limit

	if in.Status == 1 || in.Status == -1 {
		db = db.Where("status = ?", in.Status)
	}

	if len(in.Name) > 0 {
		db = db.Where("username like ? OR nickname like ? OR email like ? OR phone like ?", "%"+in.Name+"%", "%"+in.Name+"%", "%"+in.Name+"%", "%"+in.Name+"%")
	}

	var list []*proto.User
	db.Offset(int(offset)).Limit(int(in.Limit)).Order("id DESC").Find(&list)

	var count int64
	db.Model(&proto.User{}).Count(&count)

	r.List = list
	r.Count = count

	return r, nil
}

// UpdatePassword UpdatePassword
func (m *Mysql) UpdatePassword(info *proto.IDPassword) (rsp *proto.User, err error) {
	rsp, err = m.InfoByID(&proto.ID{Id: info.Id})
	if err != nil {
		return nil, err
	}

	rsp.Salt = random.String(6)
	rsp.Password = utils.HashPassword(info.Password, rsp.Salt)

	err = model.UseDB().Save(rsp).Error
	if err != nil {
		return nil, err
	}

	return rsp, nil
}
