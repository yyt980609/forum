package redis

import (
	"context"
	"fmt"
	"forum/settings"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		Username: cfg.Username,
	})
	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}
