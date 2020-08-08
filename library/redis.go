package library

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func InitRedis()  {
	Redis = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Address,
		Password: Config.Redis.Password, // no password set
		DB:       Config.Redis.DB,  // use default DB
	})

	_, err := Redis.Ping(context.TODO()).Result()
	if err != nil {
		Log.Warning("Redis initialization failed.")
	} else {
		Log.Info("Redis initialized successfully.")
	}
}