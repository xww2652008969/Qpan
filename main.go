package main

import (
	"Qpan/db"
	"Qpan/model"
)

func main() {
	db.Condb.AutoMigrate(&model.User{})
}
