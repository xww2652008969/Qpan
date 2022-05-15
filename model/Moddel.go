package model

import (
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
type File struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	name        string
	Md5         string
	Fileaddress string
}
