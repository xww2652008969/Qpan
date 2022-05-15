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
	password  string
	Name      string
	Avatar    byte
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
