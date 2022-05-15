package main

import (
	"Qpan/db"
	"Qpan/model"
	"fmt"
)

func main() {
	fmt.Println(db.Condb)
	db.Condb.AutoMigrate(&model.User{})
}
