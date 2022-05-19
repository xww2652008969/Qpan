package system

import (
	"Qpan/model"
	"Qpan/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c gin.Context) {
	var r model.User_r
	_ = c.ShouldBindJSON(&r)
	//需要鉴权未完成
	user := &model.User{Name: r.Name, Password: r.Passwd}
	err, userreun := services.Register(user)
	fmt.Print(err, userreun)
	//需要完成出参表
}
