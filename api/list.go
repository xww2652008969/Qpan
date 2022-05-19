package v1

import "Qpan/api/v1/system"

type ApiGroup struct {
	Systemapi system.Systemapi
}

var ApiGroupApp = new(ApiGroup)
