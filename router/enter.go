package router

import (
	"github.com/Joword/chatbgi-manager/router/example"
	"github.com/Joword/chatbgi-manager/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
