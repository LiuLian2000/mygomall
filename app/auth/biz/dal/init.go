package dal

import (
	"github.com/Group-lifelong-youth-training/mygomall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	// mysql.Init()
}

func Shutdown() {
	// sqlDB, _ := mysql.DB.DB()
	// sqlDB.Close()
	redis.RedisClient.Close()
}
