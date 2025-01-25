package model

import (
	"context"
	"math/rand"
	"time"

	"github.com/Group-lifelong-youth-training/mygomall/app/auth/biz/dal/redis"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/patrickmn/go-cache"
	redisp "github.com/redis/go-redis/v9"
)

const redisRandomScope = 1

var prefix = "MYGOMALL:"

var U2TPrefix = prefix + "U2T:"

var T2UPrefix = prefix + "T2U:"

var c = cache.New(5*time.Minute, 10*time.Minute)

// Get 读取字符串缓存, 其中找到了返回 True，没找到返回 False，异常也返回 False
func Get(ctx context.Context, key string) (value string, ok bool, err error) {

	if x, found := c.Get(key); found {
		return x.(string), true, nil
	}

	//缓存没有命中，Fallback 到 Redis
	cachedResult := redis.RedisClient.Get(ctx, key)

	value, err = cachedResult.Result()

	switch {
	case err == redisp.Nil:
		return "", false, nil
	case err != nil:
		klog.Errorf("Cached Redis ERR: %v", err)
		return "", false, err
	default:
		c.Set(key, value, cache.DefaultExpiration)
		return value, true, nil
	}

}

// 如果登录多次，或返回相同的token
// GetWithFunc 从缓存中获取字符串，如果不存在调用 Func 函数获取
func GetWithFunc(ctx context.Context, key string, f func(ctx context.Context, key string) (string, error)) (string, error) {

	value, ok, err := Get(ctx, key)
	if err != nil {
		return "", err
	}

	if ok {
		return value, nil
	}

	// 如果不存在，那么就获取它
	value, err = f(ctx, key)

	if err != nil {
		return "", err
	}

	Write(ctx, key, value)

	return value, nil
}

// Write 写入本地缓存和redis
func Write(ctx context.Context, key string, value string) {

	//TODO 本地成功，写入redis失败有没有问题
	//多个认证服务，一个服务端写入本地成功，写入Redis失败，这时同一个用户再另一个服务端申请token，一个用户有两个token
	c.Set(key, value, cache.DefaultExpiration)

	err := redis.RedisClient.Set(ctx, key, value, 120*time.Hour+time.Duration(rand.Intn(redisRandomScope))*time.Second).Err()
	if err != nil {
		klog.Errorf("Redis Error: %v", err)
	}

}
