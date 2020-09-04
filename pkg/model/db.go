package model

import (
	"fmt"

	"gorm.io/driver/mysql"

	"github.com/isfk/aio/config"
	log "github.com/micro/go-micro/v2/logger"
	"gorm.io/gorm"
)

// DBDrivers DBDrivers
var DBDrivers map[string]*gorm.DB

// DBInit DBInit
func DBInit(configs map[string]config.DBConf) {
	DBDrivers = make(map[string]*gorm.DB)

	for name, conf := range configs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{
			//DisableForeignKeyConstraintWhenMigrating: true,
		})

		if err != nil {
			log.Errorf("db [%s] connect fail. err: %v", name, err)
			continue
		}
		log.Infof("db [%s] connected.", name)
		DBDrivers[name] = db
	}
}

// UseDB UseDB
func UseDB(name ...string) *gorm.DB {
	k := "default"
	if len(name) > 0 {
		k = name[0]
	}

	if _, ok := DBDrivers[k]; ok {
		return DBDrivers[k]
	}

	log.Errorf("db [%s] not exist.", k)
	return nil
}
