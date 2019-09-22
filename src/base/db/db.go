package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var DataEngine *xorm.Engine

// 初始化数据库
func InitDB(path string, maxIdel, maxOpen int) {
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
	engine.SetMaxIdleConns(maxIdel)
	engine.SetMaxOpenConns(maxOpen)
	engine.SetColumnMapper(core.SameMapper{})
	DataEngine = engine
}
