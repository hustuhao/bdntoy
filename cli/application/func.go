package application

import (
	"turato.com/bdntoy/config"
	"turato.com/bdntoy/service"
)

func getService() *service.Service {
	return config.Instance.ActiveService()
}
