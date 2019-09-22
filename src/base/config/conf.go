package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// 全局配置索引
var GlobalConfig *config

// 初始化读取配置文件
func InitConfig(path string) {
	configVip := viper.New()
	configVip.SetConfigFile(path)

	// 读取配置
	if err := configVip.ReadInConfig(); err != nil {
		panic(err)
	}

	// 配置映射到结构体
	GlobalConfig = &config{}
	if err := configVip.Unmarshal(GlobalConfig); err != nil {
		panic(err)
	}

	// 这里可以做检查，如果配置文件相关配置项异常亦可以不启动
	fmt.Printf("当前读取到配置文件：\n%+v\n", *GlobalConfig)
}

type config struct {
	Server serverConfig `mapstructure:"server_config"`
	Logger loggerConfig `mapstructure:"logger_config"`
	Mysql  mysqlConfig  `mapstructure:"mysql_config"`
	Redis  redisConfig  `mapstructure:"redis_config"`
}

type serverConfig struct {
	HttpPort string `mapstructure:"http_port"`
}

type loggerConfig struct {
	Level        string        `mapstructure:"level"`
	Path         string        `mapstructure:"path"`
	MaxAge       time.Duration `mapstructure:"max_age"`
	RotationTime time.Duration `mapstructure:"rotation_time"`
}

type mysqlConfig struct {
	Connection string `mapstructure:"connection"`
	MaxIdle    int    `mapstructure:"max_idel"`
	MaxOpen    int    `mapstructure:"max_open"`
}

type redisConfig struct {
	Connection string `mapstructure:"connection"`
}
