package main

import (
	_ "github.com/go-sql-driver/mysql"

	_ "cmd/docs"

	"base/config"
	"base/db"
	"base/logger"
	"base/router"
	"base/validator"
)

// @title golang-web-template 接口文档
// @version 1.0
// @host localhost:8080
// @BasePath /gwt/api/v1
func main() {
	// 读取配置文件
	config.InitConfig("./doc/default-conf.yml")

	all := config.All
	l := all.Logger
	m := all.Mysql
	s := all.Server

	// 初始化日志
	logger.InitLogger(l.Level, "golang-web-template", l.Path, l.MaxAge, l.RotationTime)
	logger.Debug("debug日志")
	logger.Info("info日志")
	logger.Warn("warn日志")
	logger.Error("error日志")

	// 初始化数据库
	db.InitDB(m.Connection, m.MaxIdle, m.MaxOpen)

	// 初始化redis ...

	// 初始化验证器
	validator.InitValidator()

	// 初始化pprof ...

	// 初始化路由
	router.InitRouter(s.HttpPort)
}
