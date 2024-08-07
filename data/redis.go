package data

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/yosa12978/hjkl/config"
)

var (
	rdb     redis.Cmdable
	rdbOnce sync.Once
)

func Redis(ctx context.Context) redis.Cmdable {
	rdbOnce.Do(func() {
		cfg := config.Read()
		rdb = redis.NewClient(&redis.Options{
			Addr:     cfg.Redis.Addr,
			DB:       cfg.Redis.Db,
			Password: cfg.Redis.Password,
		})
		if err := rdb.Ping(ctx).Err(); err != nil {
			panic(err)
		}
	})
	return rdb
}
