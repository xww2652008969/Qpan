package model

import (
	"Qpan/db"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Password  string
	Name      string
	Avatar    byte `gorm:"type:MEDIUMBLOB"`
	Email     *string
	QQid      int
	Wechatid  *string
	Active    bool
	Filespace string
}

//创建用户如果成功返回ture否则返回false
func (user User) Create() bool {
	result := db.Condb.Create(&user)
	if result.Error != nil {
		return true
	} else {
		return false
	}
}

//查询信息如果成功返回ture和User 否则返回空USer和false
func (user User) Getuser(m map[string]interface{}) (User, bool) {
	var cuser User
	db.Condb.Where(m).First(&cuser)
	if db.Condb.Error == nil {
		return cuser, true
	} else {
		return cuser, false
	}
}

// 修改user表 如果成功则返回ture和影响行数,否则返回false，和影响行数
func (user User) Updateuser(m map[string]interface{}, o map[string]string) (bool, int64) {
	ref := db.Condb.Model(&user).Where(m).Update(o["key"], o["value"])
	fmt.Println(ref.RowsAffected)
	if db.Condb.Error == nil {
		return true, ref.RowsAffected
	} else {
		return false, ref.RowsAffected
	}
}
