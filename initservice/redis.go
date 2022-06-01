package initservice

import (
	"Qpan/global"
	"github.com/go-redis/redis/v8"
)

func Initredis() {
	global.QP_redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
