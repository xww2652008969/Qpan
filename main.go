package main

import (
	"Qpan/db"
	"Qpan/model"
	"fmt"
)

func main() {
	db.Condb.AutoMigrate(&model.User{})
	var user model.User
	db.Condb.Find(&user, model.User{Name: "xww"})
	fmt.Println(db.Condb.Error)
}
