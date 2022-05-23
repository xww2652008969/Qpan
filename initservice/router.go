package initservice

import (
	routert "Qpan/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := routert.RouterGroupApp.System
	PrivateGroup := Router.Group("")
	{
		systemRouter.Inituserrouter(PrivateGroup)
		systemRouter.InitFilerouter(PrivateGroup)
	}
	return Router
}
