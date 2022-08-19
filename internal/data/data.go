package data

import (
	"time"

	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/golang-web-template/internal/base/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/google/wire"
	"xorm.io/core"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo)

// Data data
type Data struct {
	db *xorm.Engine
}

// NewData new data instance
func NewData(c *conf.Data, l logger.Logger, db *xorm.Engine) (*Data, func(), error) {
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

// NewDB new database instance
func NewDB(c *conf.Data) *xorm.Engine {
	// 创建数据库连接
	engine, err := xorm.NewEngine("mysql", c.Database.Connection)
	if err != nil {
		panic(err)
	}

	// 测试数据库连接
	if err = engine.Ping(); err != nil {
		panic(err)
	}

	// 配置数据库相关信息
	if c.Database.MaxIdleConn > 0 {
		engine.SetMaxIdleConns(c.Database.MaxIdleConn)
	}
	if c.Database.MaxOpenConn > 0 {
		engine.SetMaxOpenConns(c.Database.MaxOpenConn)
	}
	if c.Database.ConnMaxLifeTime > 0 {
		engine.SetConnMaxLifetime(time.Duration(c.Database.ConnMaxLifeTime) * time.Second)
	}
	engine.SetColumnMapper(core.GonicMapper{})
	return engine
}
