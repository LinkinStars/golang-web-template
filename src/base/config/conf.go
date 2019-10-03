package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// All 全部配置索引
var All *AllConfig

// InitConfig 初始化读取配置文件
func InitConfig(path string) {
	configVip := viper.New()
	configVip.SetConfigFile(path)

	// 读取配置
	if err := configVip.ReadInConfig(); err != nil {
		panic(err)
	}

	// 配置映射到结构体
	All = &AllConfig{}
	if err := configVip.Unmarshal(All); err != nil {
		panic(err)
	}

	// 这里可以做检查，如果配置文件相关配置项异常亦可以不启动
	fmt.Printf("当前读取到配置文件：\n%+v\n", *All)
}

// AllConfig 全部配置文件
type AllConfig struct {
	Server ServerConfig `mapstructure:"server_config"`
	Logger LoggerConfig `mapstructure:"logger_config"`
	Mysql  MysqlConfig  `mapstructure:"mysql_config"`
	Redis  RedisConfig  `mapstructure:"redis_config"`
}

// ServerConfig 服务配置
type ServerConfig struct {
	HttpPort string `mapstructure:"http_port"`
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level        string        `mapstructure:"level"`
	Path         string        `mapstructure:"path"`
	MaxAge       time.Duration `mapstructure:"max_age"`
	RotationTime time.Duration `mapstructure:"rotation_time"`
}

// MysqlConfig 数据库配置
type MysqlConfig struct {
	Connection string `mapstructure:"connection"`
	MaxIdle    int    `mapstructure:"max_idel"`
	MaxOpen    int    `mapstructure:"max_open"`
}

// RedisConfig 缓存配置
type RedisConfig struct {
	Connection string `mapstructure:"connection"`
}
