package util

import (
	"fmt"
	"time"
)

/**
 * @ClassName redisStore
 * @Description TODO
 * @Author khr
 * @Date 2023/5/9 15:02
 * @Version 1.0
 */

const CAPTCHA = "captcha:"

type RedisStore struct {
}

// 实现设置 captcha 的方法
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := Redis.Set(ctx, key, value, time.Minute*2).Err()
	return err
}

// 实现获取 captcha 的方法
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	//获取 captcha
	val, err := Redis.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//如果clear == true, 则删除
	if clear {
		err := Redis.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// 实现验证 captcha 的方法
func (r RedisStore) Verify(id, answer string, clear bool) bool {

	fmt.Println("进入自己写的方法")

	v := r.Get(id, clear)
	fmt.Println(v, "v")
	return v == answer
}
