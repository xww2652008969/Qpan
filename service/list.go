package service

import "Qpan/service/system"

type ServiceGroup struct {
	Systemservice system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
