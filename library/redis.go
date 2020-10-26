package library

import (
	"context"
	Package_Redis "github.com/go-redis/redis/v8"
	"time"
)

var Instance_Redis *Package_Redis.Client

func init_Redis()  {
	//	If the cache is not turned on, skip redis initialization
	//	如果没有开启缓存则跳过redis初始化
	if !Config.App.Cache {
		return
	}

	Instance_Redis = Package_Redis.NewClient(&Package_Redis.Options{
		Addr:     Config.Cache.Address,
		Password: Config.Cache.Password, // no password set
		DB:       Config.Cache.DB,  // use default DB
	})

	_, err := Instance_Redis.Ping(context.TODO()).Result()
	if err != nil {
		Instance_Logrus.Warning("redis initialization failed.")
	} else {
		Instance_Logrus.Info("redis initialized successfully.")
	}
}


type redis struct {}

func NewCache() *redis {
	return &redis{}
}

// Get String
// 获取字符串
func (this *redis) GetString (key string) (string, error) {
	return Instance_Redis.Get(context.Background(), key).Result()
}


// Set String
// 设置字符串
func (this *redis) SetString (key string, value string, time time.Duration) {
	_ = Instance_Redis.Set(context.Background(), key, value, time).Err()
	return
}


// Delete String
// 删除字符串
func (this *redis) DeleteString (list ...string) {
	_ = Instance_Redis.Del(context.Background(), list...).Err()
	return
}


// Get Hash
// 获取哈希
func (this *redis) GetHash (key string, field string) (string, error) {
	return Instance_Redis.HGet(context.Background(), key, field).Result()
}


// Set Hash
// 设置哈希
func (this *redis) SetHash (key string, value map[string]string) {
	var tmp = make(map[string]interface{})
	for k, v := range value {
		tmp[k] = v
	}
	_ = Instance_Redis.HSet(context.Background(), key, tmp).Err()
	return
}


// Delete Hash
// 删除哈希
func (this *redis) DeleteHash (key string, list ...string) {
	_ = Instance_Redis.HDel(context.Background(), key, list...).Err()
	return
}


// Clear all caches in the current DB
// 清除当前DB内所有缓存
func (this *redis) ClearAllCache() {
	Instance_Redis.FlushDB(context.Background())
	return
}



