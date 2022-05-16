package main

import (
	"Qpan/db"
	"Qpan/model"
)

func main() {
	db.Condb.AutoMigrate(&model.User{})
	user := map[string]interface{}{"id": "61"}
	gat := map[string]string{
		"key":   "id",
		"value": "6223",
	}
	model.User.Updateuser(model.User{}, user, gat)
}
