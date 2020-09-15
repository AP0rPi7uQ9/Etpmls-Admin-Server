package library

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var RedisClient *redis.Client


func InitRedis()  {
	//	If the cache is not turned on, skip redis initialization
	//	如果没有开启缓存则跳过redis初始化
	if !Config.App.Cache {
		return
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     Config.Cache.Address,
		Password: Config.Cache.Password, // no password set
		DB:       Config.Cache.DB,  // use default DB
	})

	_, err := RedisClient.Ping(context.TODO()).Result()
	if err != nil {
		Log.Warning("Redis initialization failed.")
	} else {
		Log.Info("Redis initialized successfully.")
	}
}


type Redis struct {}


// Get String
// 获取字符串
func (this *Redis) GetString (key string) (string, error) {
	return RedisClient.Get(context.Background(), key).Result()
}


// Set String
// 设置字符串
func (this *Redis) SetString (key string, value string, time time.Duration) {
	_ = RedisClient.Set(context.Background(), key, value, time).Err()
	return
}


// Delete String
// 删除字符串
func (this *Redis) DeleteString (list ...string) {
	_ = RedisClient.Del(context.Background(), list...).Err()
	return
}


// Get Hash
// 获取哈希
func (this *Redis) GetHash (key string, field string) (string, error) {
	return RedisClient.HGet(context.Background(), key, field).Result()
}


// Set Hash
// 设置哈希
func (this *Redis) SetHash (key string, value map[string]string) {
	_ = RedisClient.HSet(context.Background(), key, value).Err()
	return
}


// Delete Hash
// 删除哈希
func (this *Redis) DeleteHash (key string, list ...string) {
	_ = RedisClient.HDel(context.Background(), key, list...).Err()
	return
}





