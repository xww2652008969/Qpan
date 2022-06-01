package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Uuid      uuid.UUID
	Password  string
	Name      string
	Email     *string
	Filespace string
}
