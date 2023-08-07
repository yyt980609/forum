package mysql

import (
	"fmt"
	"forum/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(cfg *config.MySQLConfig) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn}), &gorm.Config{})
	return db, err
}
