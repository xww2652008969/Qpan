package system

import (
	"Qpan/api"
	"github.com/gin-gonic/gin"
)

type Userrouter struct {
}

func (u *Userrouter) Inituserrouter(Router *gin.RouterGroup) {
	Api := v1.ApiGroupApp.Systemapi.Verifyapi
	user := Router.Group("/user")
	{
		user.POST("/register", Api.Register)
	}
}
