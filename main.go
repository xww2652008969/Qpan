package main

import (
	"Qpan/db"
	"Qpan/model"
)

func main() {
	db.Condb.AutoMigrate(&model.User{})
	user := model.User{Name: "xww"}
	db.Create(user)
}
