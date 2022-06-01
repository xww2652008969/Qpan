package utils

import (
	"Qpan/global"
	"context"
)

var ctx = context.Background()

func Addredis(key string, values ...interface{}) (int64, error) {
	return global.QP_redis.HSet(global.QP_ctx, key, values).Result()
}
func Getredis(key string, feild string) (string, error) {
	return global.QP_redis.HGet(global.QP_ctx, key, feild).Result()
}
func Delredis(key string, feild string) (int64, error) {
	return global.QP_redis.HDel(global.QP_ctx, key, feild).Result()
}
