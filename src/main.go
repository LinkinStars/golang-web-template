package main

import (
	"base/config"
	"base/db"
	"base/logger"
	"base/router"
	"base/validator"
)

func main() {
	// 读取配置文件
	config.InitConfig("./default-conf.yml")

	c := config.GlobalConfig
	l := c.Logger
	m := c.Mysql
	s := c.Server

	// 初始化日志
	logger.InitLogger("golang-web-template", l.Path, l.MaxAge, l.RotationTime, l.Level)
	logger.Debug("debug日志")
	logger.Info("info日志")
	logger.Warn("warn日志")
	logger.Error("error日志")

	// 初始化数据库
	db.InitDB(m.Connection, m.MaxIdle, m.MaxOpen)

	// 初始化redis

	// 初始化验证器
	validator.InitValidator()

	// 初始化pprof

	// 初始化路由
	router.InitRouter(s.HttpPort)
}
