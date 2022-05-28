package system

import (
	"Qpan/model"
	"Qpan/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

//用户注册

type Verifyapi struct {
}

func (v *Verifyapi) Register(c *gin.Context) {
	var r model.UserR
	_ = c.ShouldBindJSON(&r)
	//需要检验传入参数
	user := &model.User{Name: r.Name, Password: r.Passwd}
	err, userreun := userService.Register(user)
	fmt.Print(err, userreun)
	//需要完成出参表(不是怎么详细)
	if err != nil {
		model.Errorres(nil, "失败", c)
	} else {
		model.Successres(userreun, "成功", c)
	}
}

//用户登录

func (v *Verifyapi) Login(c *gin.Context) {
	var r model.UserR
	_ = c.ShouldBindJSON(&r)
	user := &model.User{Name: r.Name, Password: r.Passwd}
	err, passwderror := userService.Login(user)
	if err == nil {
		t, _ := utils.Createtoken(passwderror.Uuid)
		tok := model.Usertoken{Token: t}
		model.Successres(tok, "登录成功", c)
	} else {
		model.Errorres(nil, "登录失败", c)
	}
}
