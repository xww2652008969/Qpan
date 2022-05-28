package model

import (
	"Qpan/global"
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Uuid      uuid.UUID
	Password  string
	Name      string
	Email     *string
	QQid      int
	Wechatid  *string
	Filespace string
}
type Usertoken struct {
	Token string
}

// Create 创建用户如果成功返回ture否则返回false
func (user User) Create() bool {
	result := global.QP_db.Create(&user)
	if result.Error != nil {
		return true
	} else {
		return false
	}
}

// Getuser 查询信息如果成功返回ture和User 否则返回空USer和false
func (user User) Getuser(m map[string]interface{}) (User, bool) {
	var cuser User
	global.QP_db.Where(m).First(&cuser)
	if global.QP_db.Error == nil {
		return cuser, true
	} else {
		return cuser, false
	}
}

// Updateuser 修改user表 如果成功则返回ture和影响行数,否则返回false，和影响行数
func (user User) Updateuser(m map[string]interface{}, o map[string]string) (bool, int64) {
	ref := global.QP_db.Model(&user).Where(m).Update(o["key"], o["value"])
	fmt.Println(ref.RowsAffected)
	if global.QP_db.Error == nil {
		return true, ref.RowsAffected
	} else {
		return false, ref.RowsAffected
	}
}

////Deleteuser 删除user表，如果成功啧返回true和影响行数，否则返回false，同时无变化
///**
//未测试
//*/
//func (user User) Deleteuser(m map[string]interface{}, o map[string]string) (bool, int64) {
//
//}
