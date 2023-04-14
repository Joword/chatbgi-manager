package v1

import (
	"github.com/Joword/chatbgi-manager/api/v1/example"
	"github.com/Joword/chatbgi-manager/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
