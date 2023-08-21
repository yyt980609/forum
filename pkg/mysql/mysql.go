package mysql

import (
	"fmt"
	"forum/config"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init 初始化Gorm
func Init(cfg *config.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	dialector := mysql.New(
		mysql.Config{
			DriverName: "mysql",
			DSN:        dsn})
	db, err = gorm.Open(dialector, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	return err
}
func GetDB() *gorm.DB {
	return db
}
