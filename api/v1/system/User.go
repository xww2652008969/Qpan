package system

import (
	"Qpan/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

//用户注册

type Verifyapi struct {
}

func (v *Verifyapi) Register(c *gin.Context) {
	var r model.User_r
	_ = c.ShouldBindJSON(&r)
	//需要鉴权未完成
	user := &model.User{Name: r.Name, Password: r.Passwd}
	err, userreun := userService.Register(user)
	fmt.Print(err, userreun)
	//需要完成出参表(不是怎么详细)
	if err != nil {
		model.FailWithMessage(r, "完成", c)
	}
}

//用户登录

func (v *Verifyapi) Login(c *gin.Context) {
	var r model.User_r
	_ = c.ShouldBindJSON(&r)
	user := &model.User{Name: r.Name, Password: r.Passwd}
	err, passwderror := userService.Login(user)
	fmt.Print(err, passwderror)
	if err == nil {
		model.FailWithMessage(r, "登录成功", c)
	} else {
		model.FailWithMessage(r, "登录登录失败", c)
	}
}
