package service

import (
	"github.com/Joword/chatbgi-manager/service/example"
	"github.com/Joword/chatbgi-manager/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
