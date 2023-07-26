package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	Redis *redis.Client
)

func Redisinit() {

	Redis = redis.NewClient(&redis.Options{
		Addr: RedisConfig.Host + ":" + RedisConfig.Port,
		//Username:   redisCon.UserName,
		//Password:   redisCon.PassWord,
		DB:         RedisConfig.Db,
		PoolSize:   RedisConfig.PoolSize,
		MaxRetries: RedisConfig.MaxRetries,
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("redis connect get failed.%v", err)
		return
	}

	log.Printf("redis 初始化连接成功")
}
