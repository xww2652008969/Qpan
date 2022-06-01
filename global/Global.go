package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	//全局变量
	QP_db    *gorm.DB           //数据库操作
	JwtKey   = []byte("123456") //key密钥
	QP_redis *redis.Client      //redis操作
	QP_ctx   = context.Background()
)
