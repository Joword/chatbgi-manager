package core

import (
	"fmt"
	"time"

	"github.com/Joword/chatbgi-manager/global"
	"github.com/Joword/chatbgi-manager/initialize"
	"github.com/Joword/chatbgi-manager/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {

	version := "v1.0.0"

	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 ChatBGI-manager
	当前版本: `+version+`
	默认自动化文档地址:http://127.0.0.1/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:9527
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
