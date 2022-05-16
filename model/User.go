package model

import (
	"Qpan/db"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Password  string
	Name      string  `gorm:"unique"`
	Avatar    byte    `gorm:"type:MEDIUMBLOB"`
	Email     *string `gorm:"unique"`
	QQid      int     `gorm:"unique"`
	Wechatid  *string `gorm:"unique"`
	Active    bool
	Filespace string `gorm:"unique"`
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
