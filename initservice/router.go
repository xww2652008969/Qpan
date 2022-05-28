package initservice

import (
	"Qpan/middleware"
	routert "Qpan/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors())
	systemRouter := routert.RouterGroupApp.System
	PrivateGroup := Router.Group("")
	{
		systemRouter.Inituserrouter(PrivateGroup)
		systemRouter.InitFilerouter(PrivateGroup)
	}
	return Router
}
