package system

import "Qpan/service"

type Systemapi struct {
	Verifyapi //用于登录验证api
}

var (
	userService = service.ServiceGroupApp.Systemservice
)
