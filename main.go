package main

import (
	"github.com/childelins/go-skeleton/bootstrap"
	_ "github.com/childelins/go-skeleton/config"
	"github.com/childelins/go-skeleton/pkg/config"
	"github.com/childelins/go-skeleton/pkg/database"
)

func init() {
	// 配置初始化，依赖命令行 --env 参数
	config.InitConfig("")

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 初始化数据库
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()

	// 初始化缓存
	bootstrap.SetupCache()
}

func main() {
	app := BuildApp(database.DB)
	app.Run()
}
