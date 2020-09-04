package config

import (
	"fmt"

	"github.com/go-playground/log/v7"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

// Init Init
func Init(configFile string) {
	fileSource := file.NewSource(
		file.WithPath(configFile),
	)

	conf, err := config.NewConfig()
	if err != nil {
		log.WithField("err", err).Panic("config.NewConfig fail.")
		panic(err)
	}

	err = conf.Load(fileSource)
	if err != nil {
		log.WithField("err", err).Panic("config load fail.")
		panic(err)
	}

	if err := conf.Get("global").Scan(&Conf); err != nil {
		fmt.Println("err", err)
		log.WithField("err", err).Panic("config scan fail.")
		panic(err)
	}
	fmt.Println("Conf", Conf)
}

// Conf Conf
var Conf = &Config{
	DB:    map[string]DBConf{},
	Redis: map[string]RedisConf{},
}

// Config Config
type Config struct {
	DB    map[string]DBConf
	Redis map[string]RedisConf
}

// DBConf DBConf
type DBConf struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Db       string `json:"db"`
	Port     string `json:"port"`
}

// RedisConf RedisConf
type RedisConf struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     string `json:"port"`
}
