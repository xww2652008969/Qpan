package db

import "Qpan/model"

//User 创建记录
func Create(data model.User) {
	Condb.Create(&data)
}
