package model

import (
	"gorm.io/gorm"
	"time"
)

type File struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	name        string
	Md5         string
	Fileaddress string
}
