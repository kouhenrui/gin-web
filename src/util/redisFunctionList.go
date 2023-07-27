package util

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-web/src/global"
	"time"
)

var (
	ctx   = context.Background()
	Redis = global.Redis
)

// json格式化数据
func Marshal(user interface{}) []byte {
	ub, _ := json.Marshal(user)
	return ub
}
func UnMarshal(r []byte, res interface{}) (bool, interface{}) {
	err := json.Unmarshal(r, &res)
	if err != nil {
		return false, REDIS_INFORMATION_ERROR
	}
	return true, res
}

// 添加数据
func SetRedis(key string, value []byte, t int) bool {
	expire := time.Duration(t) * global.DayTime

	if err := Redis.Set(ctx, key, value, expire).Err(); err != nil {
		//log.Println(err, "redis存放错误")
		fmt.Println(err, "redis存放错误")
		return false
	}
	fmt.Println("redis存放时间", expire)
	return true
}

// set 中是否存在某个成员
func ExistRedis(key string) bool {
	a, err := Redis.Exists(ctx, key).Result()

	//fmt.Println(a, "cunchudezhi ")
	if err != nil {
		fmt.Println(err)
		return false
	}
	if a != 1 {
		return false
	}
	return true
}

// 获取数据
func GetRedis(key string) string {
	//fmt.Println(key, "建")
	result, err := Redis.Get(ctx, key).Result()
	//fmt.Println(result, "redis存储的值")
	if err != nil {
		fmt.Println("获取redis缓存错误", err)
		return ""
	}
	return result
}

// 删除数据
func DelRedis(key string) error {
	_, err := Redis.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 延长过期时间
func ExpireRedis(key string, t int) error {
	expire := time.Duration(t) * time.Second
	if err := Redis.Expire(ctx, key, expire).Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
