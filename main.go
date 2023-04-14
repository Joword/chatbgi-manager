package main

import (
	"go.uber.org/zap"

	"github.com/Joword/chatbgi-manager/core"
	"github.com/Joword/chatbgi-manager/global"
	"github.com/Joword/chatbgi-manager/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	global.GVA_VP = core.Viper()
	initialize.OtherInit()
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm()
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables()
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
