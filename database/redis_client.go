package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/web_demo/v2/config"
	"github.com/web_demo/v2/log"
)

var ctx = context.Background()

var (
	RedisCli *redis.Client
)

func CreateRedis() {
	log.Sugar.Infow("CreateRedis ", "redis_address", config.Config.GetString("redis.address"))
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     config.Config.GetString("redis.address"),
		Password: config.Config.GetString("redis.password"), // no password set
		DB:       config.Config.GetInt("redis.db"),          // use default DB
	})
	pong, err := RedisCli.Ping(ctx).Result()
	if err != nil {
		log.Sugar.Fatalw("Redis ping error.",
			"error", err.Error(),
			"pong", pong)
	}
}
