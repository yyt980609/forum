package main

import (
	"fmt"
	"forum/dao/mysql"
	"forum/dao/redis"
	"forum/logger"
	"forum/routes"
	"forum/settings"

	"go.uber.org/zap"
)

func main() {
	// 读取配置文件
	if err := settings.Init(); err != nil {
		zap.L().Error("Init config failed.", zap.String("err", err.Error()))
	}
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		zap.L().Error("Init logger failed.", zap.String("err", err.Error()))
	}
	// 将日志加载之前生成的日志追加到日志文件中
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			zap.L().Error("Sync logger failed.", zap.String("err", err.Error()))
		}
	}(zap.L())
	// mysql
	if _, err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		zap.L().Error("Init mysql failed.", zap.String("err", err.Error()))
	}
	// redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		zap.L().Error("Init redis failed.", zap.String("err", err.Error()))
	}
	defer redis.Close()
	// 路由
	r := routes.SetUp()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		zap.L().Error("Server start failed.", zap.String("err", err.Error()))
		return
	}
}
