package model

import (
	"Qpan/db"
	"errors"
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
	Name      string  `gorm:"unique"`
	Avatar    byte    `gorm:"type:MEDIUMBLOB"`
	Email     *string `gorm:"unique"`
	QQid      int     `gorm:"unique"`
	Wechatid  *string `gorm:"unique"`
	Active    bool
	Filespace string `gorm:"unique"`
}

func (user User) Create() bool {
	result := db.Condb.Create(&user)
	fmt.Println(result.Error)
	errors.Is(result.Error, gorm.ErrDryRunModeUnsupported)
	return true
}
