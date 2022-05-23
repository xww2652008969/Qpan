package service

import "Qpan/service/system"

type ServiceGroup struct {
	Systemservice system.ServiceGroup
	Filiservice   system.Fileservice
}

var ServiceGroupApp = new(ServiceGroup)
