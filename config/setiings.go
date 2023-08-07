package config

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int16  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"dbname"`
	Port     int    `mapstructure:"port"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func InitSettings() (err error) {
	// 配置文件名称(无扩展名)
	viper.SetConfigName("config")
	// 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.SetConfigType("yaml")
	// 在工作目录中查找配置
	viper.AddConfigPath("./config")
	// 查找并读取配置文件
	err = viper.ReadInConfig()
	// 处理读取配置文件的错误
	if err != nil {
		zap.L().Error("Viper read config failed.", zap.String("err", err.Error()))
		return
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		zap.L().Error("Viper unmarshal failed.", zap.String("err", err.Error()))
		return
	}
	// 热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config update")
		err = viper.Unmarshal(Conf)
		if err != nil {
			zap.L().Error("Viper unmarshal failed.", zap.String("err", err.Error()))
			return
		}
	})
	return
}
