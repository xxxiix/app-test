package main

import (
	"fmt"
	"main/controllers"
	"main/dao/mysql"
	"main/logger"
	"main/pkg/gfsd"
	"main/pkg/snowflake"
	"main/routes"
	"main/settings"

	"go.uber.org/zap"
)

func main() {
	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err = %v\n", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err = %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("init ogger success...")

	// 3. 搭建 Mysql 连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed = %v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 搭建 Redis 连接
	// 方法同上，就不写了

	// 雪花算法初始化
	if err := snowflake.Init(); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 初始化gin内置翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed = %v\n", err)
		return
	}
	// 5. 注册路由
	r := routes.Setup()

	// 6. 启动服务(同时实现优雅关机)
	gfsd.Graceful_Shutdown(r)
}
