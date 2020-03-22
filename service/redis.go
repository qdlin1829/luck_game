package service

import (
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

func init() {
	initRedis()
}

// 初始化连接
func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 20,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
