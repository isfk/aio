package casbin

import (
	"fmt"

	redisWatcher "github.com/billcobbler/casbin-redis-watcher/v2"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/isfk/aio/config"
	log "github.com/micro/go-micro/v2/logger"

	_ "github.com/go-sql-driver/mysql"
)

// GetEnforcer GetEnforcer
func GetEnforcer() (*casbin.Enforcer, error) {
	rsn := fmt.Sprintf("%s:%s", config.Conf.Redis["default"].Host, config.Conf.Redis["default"].Port)
	w, err := redisWatcher.NewWatcher(rsn, redisWatcher.Password(config.Conf.Redis["default"].Password))
	if err != nil {
		log.Error("redisWatcher.NewWatcher err ", err.Error())
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Conf.DB["default"].User, config.Conf.DB["default"].Password, config.Conf.DB["default"].Host, config.Conf.DB["default"].Port, config.Conf.DB["default"].Db)
	a, err := gormAdapter.NewAdapter("mysql", dsn, true)
	if err != nil {
		log.Error("gormAdapter.NewAdapter err ", err.Error())
		panic(err)
	}

	text := `
		[request_definition]
		r = sub, obj, act, id
		
		[policy_definition]
		p = sub, obj, act, id
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) && regexMatch(r.id, p.id)
		`

	m, err := model.NewModelFromString(text)
	if err != nil {
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}
	enforcer.SetWatcher(w)
	return enforcer, nil
}
