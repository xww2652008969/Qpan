package system

import "Qpan/service"

type Systemapi struct {
	Verifyapi //用于登录验证api
	Fileapi
}

var (
	userService = service.ServiceGroupApp.Systemservice
	fileservice = service.ServiceGroupApp.Filiservice
)
