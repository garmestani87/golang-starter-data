package cache

import (
	"golang-starter-data/config"
	"time"

	"github.com/go-redis/redis"
)

var rdc *redis.Client

func InitRedis(cfg *config.Config) {

	rdc = redis.NewClient(&redis.Options{
		Addr:               cfg.Cache.Redis.Uri,
		Password:           cfg.Cache.Redis.Password,
		DB:                 cfg.Cache.Redis.Db,
		ReadTimeout:        time.Duration(cfg.Cache.Redis.ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(cfg.Cache.Redis.WriteTimeout) * time.Second,
		PoolSize:           cfg.Cache.Redis.PoolSize,
		PoolTimeout:        time.Duration(cfg.Cache.Redis.PoolTimeout) * time.Second,
		IdleTimeout:        time.Duration(cfg.Cache.Redis.IdleTimeout) * time.Second,
		IdleCheckFrequency: time.Duration(cfg.Cache.Redis.IdleCheckFrequency) * time.Second,
	})
}

func GetClient() *redis.Client {
	return rdc
}

func CloseRedis() {
	err := rdc.Close()
	if err != nil {
		panic(err)
		return
	}
}
