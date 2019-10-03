package db

import (
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

// Engine 全局数据引擎
var Engine *xorm.Engine

// InitDB 初始化数据库
func InitDB(path string, maxIdle, maxOpen int) {
	// 创建数据库连接
	engine, err := xorm.NewEngine("mysql", path)
	if err != nil {
		panic(err)
	}

	// 测试数据库连接
	if err = engine.Ping(); err != nil {
		panic(err)
	}

	// 打印数据库执行语句
	engine.ShowSQL(true)

	// 配置数据库相关信息
	engine.SetMaxIdleConns(maxIdle)
	engine.SetMaxOpenConns(maxOpen)
	engine.SetColumnMapper(core.GonicMapper{})
	Engine = engine
}
