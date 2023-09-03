package dao

import (
	"forum/models"
	"forum/pkg/mysql"
	"time"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	db := mysql.GetDB()
	post.CreateTime = time.Now()
	tx := db.Table("post").Create(post)
	if tx.Error != nil {
		zap.L().Info("Insert post.")
		err = tx.Error
		return
	} else {
		return
	}
}

// SelectPostById 查看帖子详情
func SelectPostById(id int64) (data *models.Post, err error) {
	db := mysql.GetDB()
	result := db.Table("post").Where("post_id = ?", id).First(&data)
	err = result.Error
	return
}

// SelectPostList 分页查询帖子
func SelectPostList(page *models.Page[*models.Post], order string) (err error) {
	db := mysql.GetDB()
	// 统计总数
	db.Table("post").Count(&page.Total)
	// 计算分页信息，查询数据
	err = db.Table("post").Order(order).Scopes(mysql.Paginate(page)).Find(&page.Data).Error
	return
}
