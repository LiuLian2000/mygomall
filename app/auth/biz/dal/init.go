package dal

import (
	"github.com/Group-lifelong-youth-training/mygomall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	// mysql.Init()
}
