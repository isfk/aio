package model

import (
	"github.com/isfk/aio/config"
	"github.com/isfk/aio/pkg/model"
	proto "github.com/isfk/aio/srv/role/proto/role"
)

// Init Init
func Init(config *config.Config) {
	model.CacheInit(config.Redis)
	model.DBInit(config.DB)
	// autoMigrate
	autoMigrate()
}

func autoMigrate(db ...string) {
	name := "default"
	if len(db) > 0 {
		name = db[0]
	}
	model.UseDB(name).Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4").AutoMigrate(&proto.Role{})
}
