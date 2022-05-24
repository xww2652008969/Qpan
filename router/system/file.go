package system

import (
	"Qpan/api"
	"Qpan/middleware"
	"github.com/gin-gonic/gin"
)

type Filerouter struct {
}

func (u *Filerouter) InitFilerouter(Router *gin.RouterGroup) {
	Api := v1.ApiGroupApp.Systemapi.Fileapi
	user := Router.Group("/file").Use(middleware.Jwtauth())
	{
		user.POST("/upload", Api.Upload)
		user.GET("/getfile", Api.Getfile)
		user.GET("/list", Api.Getfilelist)
	}
}
