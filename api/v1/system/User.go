package system

import (
	"Qpan/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
