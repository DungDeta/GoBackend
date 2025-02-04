package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"myproject/global"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Db,
		PoolSize: r.Pool,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis connect error", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("Redis connect success")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	// Set a value
	err := global.Rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		global.Logger.Error("Set key error", zap.Error(err))
		panic(err)
	}
	res, err := global.Rdb.Get(ctx, "key").Result()
	if err != nil {
		global.Logger.Error("Get key error", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("key", zap.String("key", res))
}
