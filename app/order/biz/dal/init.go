package dal

import (
	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/dal/mysql"
	"github.com/Group-lifelong-youth-training/mygomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
