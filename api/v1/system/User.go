package system

import (
	"Qpan/model"
	"Qpan/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

//用户注册

type Userapi struct {
}

func (v *Userapi) Register(c *gin.Context) {
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

func (v *Userapi) Login(c *gin.Context) {
	var r model.UserR
	_ = c.ShouldBindJSON(&r)
	user := &model.User{Name: r.Name, Password: r.Passwd}
	err, passwderror := userService.Login(user)
	if err == nil {
		t, _ := utils.Createtoken(passwderror.Uuid)
		tok := model.Usertoken{Token: t}
		utils.Addredis(fmt.Sprintf("%v", passwderror.Uuid), "jwt", tok.Token)
		model.Successres(tok, "登录成功", c)
	} else {
		model.Errorres(nil, "登录失败", c)
	}
}

//用户更新密码
func (u Userapi) Updateuser(c *gin.Context) {
	var r model.Updateuser
	uuid := c.Request.Header.Get("uuid")
	_ = c.ShouldBindJSON(&r)
	r.Uuid = uuid
	err := userService.Updateuser(r)
	if err == nil {
		utils.Delredis(uuid, "jwt")
		model.Successres(nil, "修改密码成功", c)
	} else {
		model.Errorres(nil, "原密码错误", c)
	}
}
