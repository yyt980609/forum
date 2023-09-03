package mysql

import (
	"fmt"
	"forum/config"
	"forum/models"

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

func Paginate[T any](page *models.Page[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.CurrentPage <= 0 {
			page.CurrentPage = 0
		}
		switch {
		case page.PageSize > 100:
			page.PageSize = 100
		case page.PageSize <= 0:
			page.PageSize = 10
		}
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		p := page.CurrentPage
		if page.CurrentPage > page.Pages {
			p = page.Pages
		}
		pageSize := page.PageSize
		offset := int((p - 1) * pageSize)
		return db.Offset(offset).Limit(int(pageSize))
	}
}
