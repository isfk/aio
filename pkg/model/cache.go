package model

import (
	"fmt"

	"github.com/isfk/aio/config"

	"github.com/go-cache/cache"
	"github.com/go-redis/redis/v7"
	log "github.com/micro/go-micro/v2/logger"
)

// cacheDrivers cacheDrivers
var cacheDrivers map[string]*cache.Client

// UseCache UseCache
func UseCache(name ...string) *cache.Client {
	k := "default"
	if len(name) > 0 {
		k = name[0]
	}

	if _, ok := cacheDrivers[k]; ok {
		return cacheDrivers[k]
	}

	log.Errorf("redis [%k] not exist.", k)
	return nil
}

// CacheInit CacheInit
func CacheInit(configs map[string]config.RedisConf) {
	cacheDrivers = make(map[string]*cache.Client)

	for name, conf := range configs {
		redisConf := redis.Options{
			Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
			Password: conf.Password,
		}

		client := cache.NewClient(redisConf)
		_, err := client.RedisClient.Ping().Result()

		if err != nil {
			log.Errorf("redis [%s] connect fail.", name)
			continue
		}

		log.Infof("redis [%s] connected.", name)
		cacheDrivers[name] = client
	}
}
