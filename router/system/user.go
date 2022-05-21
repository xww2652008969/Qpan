package system

import (
	"Qpan/api"
	"Qpan/middleware"
	"github.com/gin-gonic/gin"
)

type Userrouter struct {
}

func (u *Userrouter) Inituserrouter(Router *gin.RouterGroup) {
	Api := v1.ApiGroupApp.Systemapi.Verifyapi
	user := Router.Group("/user").Use(middleware.Jwtauth())
	{
		user.POST("/register", Api.Register)
		user.POST("/login", Api.Login)
	}
}
