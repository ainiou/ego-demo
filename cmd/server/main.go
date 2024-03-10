package main

import (
	"ego-demo/internal/server/application"
	"ego-demo/pkg/econfig"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"go.uber.org/zap"
)

// 本地运行命令 export EGO_DEBUG=true && go run cmd/server/main.go  --config=configs/server/local.toml
// export EGO_DEBUG=true 让日志输出到控制台 [可选]
// --config=configs/server/local.toml 指定配置文件 【部署的时候使用EGO_CONFIG_PATH环境变量指定配置文件的路径】
func main() {
	// 注意：这一步要在ego.New之前执行
	if err := econfig.Init(); err != nil {
		// 加载配置失败，直接panic
		elog.Panic("初始化配置失败", zap.Error(err))
		return
	}
	app := ego.New()

	// 多cmd运行的时候 通个加载不同的application包来实现不同程序的启动
	if err := application.InitApp(app); err != nil {
		panic(err)
	}

	err := app.Run()
	if err != nil {
		elog.Panic("start up error: " + err.Error())
	}
}
