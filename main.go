package main

import (
	"fmt"
	"forum/config"
	"forum/router"
	"forum/utils/logger"
	"forum/utils/mysql"
	"forum/utils/redis"
	"forum/utils/snowflake"

	"go.uber.org/zap"
)

func main() {
	// 读取配置文件
	if err := config.InitSettings(); err != nil {
		zap.L().Error("Init config failed.", zap.String("err", err.Error()))
	}
	// 初始化日志组件
	if err := logger.Init(config.Conf.LogConfig); err != nil {
		zap.L().Error("Init logger failed.", zap.String("err", err.Error()))
	}
	// 将日志组件加载之前生成的日志追加到日志文件中
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			zap.L().Error("Sync logger failed.", zap.String("err", err.Error()))
		}
	}(zap.L())
	// mysql
	if _, err := mysql.Init(config.Conf.MySQLConfig); err != nil {
		zap.L().Error("Init mysql failed.", zap.String("err", err.Error()))
	}
	// redis
	if err := redis.Init(config.Conf.RedisConfig); err != nil {
		zap.L().Error("Init redis failed.", zap.String("err", err.Error()))
	}
	defer redis.Close()
	// 初始化分布式id生成器
	if err := snowflake.Init(uint16(config.Conf.MachineID), config.Conf.StartTime); err != nil {
		zap.L().Error("Init sony flake failed.", zap.String("err", err.Error()))
	}
	// 路由
	r := router.SetUp()
	err := r.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		zap.L().Error("Server start failed.", zap.String("err", err.Error()))
		return
	}
}
