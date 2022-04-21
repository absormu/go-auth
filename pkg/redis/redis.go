package redis

import (
	cm "github.com/absormu/go-auth/pkg/configuration"
	"github.com/go-redis/redis/v8"
)

func RedisDBInit(db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cm.Config.RedisDBAddr + ":" + cm.Config.RedisPort,
		Password: cm.Config.RedisDBPassword,
		DB:       db,
	})
	return rdb
}
