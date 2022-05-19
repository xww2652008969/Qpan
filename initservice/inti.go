package initservice

import (
	"Qpan/global"
	"Qpan/model"
)

func Init() {
	//注册数据库并且自动迁移数据库
	DB()
	global.QP_db.AutoMigrate(model.User{})
	global.QP_db.AutoMigrate(model.File{})
	rou := Routers()
	rou.Run(":8888")
}
