package service

import (
	"forum/common"
	"forum/dao"
	"forum/models"
	fRedis "forum/pkg/redis"
	"forum/pkg/snowflake"
	fError "forum/utils/forum_error"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (error, int64) {
	id, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("GetID with error.", zap.Error(err))
		return err, id
	}
	post.PostId = id
	err = dao.CreatePost(post)
	if err != nil {
		zap.L().Error("CreatePost with error.", zap.Error(err))
		return err, id
	}
	return initPostScore(id)
}

// 初始化帖子的分数
func initPostScore(id int64) (error, int64) {
	client := fRedis.GetRedis()
	context := fRedis.GetContext()
	pipeline := client.TxPipeline()
	// 1.在redis中插入帖子的发布时间
	pipeline.ZAdd(context, common.GetRedisKey(common.KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: id,
	})
	// 2.在redis中插入帖子的分数
	pipeline.ZAdd(context, common.GetRedisKey(common.KeyPostScoreZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: id,
	})
	_, err := pipeline.Exec(context)
	if err != nil {
		zap.L().Error("RedisTx with error.", zap.Error(err))
		return err, id
	}
	return nil, id
}

// GetPostDetail 查看帖子详情
func GetPostDetail(idStr string) (data *models.ApiPostDetail, err error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, fError.New(fError.CodeInvalidParam)
	}
	post, err := dao.SelectPostById(id)
	if err != nil {
		zap.L().Error("SelectPostById failed.", zap.Error(err))
		return nil, fError.New(fError.CodeNoData)
	}
	user, err := dao.SelectUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("SelectUserById failed.", zap.Error(err))
		return nil, fError.New(fError.CodeNoData)
	}
	community, err := dao.SelectCommunityById(post.CommunityId)
	if err != nil {
		zap.L().Error("SelectCommunityById failed.", zap.Error(err))
		return nil, fError.New(fError.CodeNoData)
	}
	data = &models.ApiPostDetail{
		AuthorName: user.Username,
		Post:       post,
		Community:  community,
	}
	return
}

// GetPostList 分页查询帖子
func GetPostList(currentPage, pageSize int, order string) (resultPage *models.Page[*models.ApiPostDetail], err error) {
	page := &models.Page[*models.Post]{
		CurrentPage: int64(currentPage),
		PageSize:    int64(pageSize),
		Data:        make([]*models.Post, 0),
	}
	err = dao.SelectPostList(page, order)
	if err != nil {
		zap.L().Error("SelectPostById failed.", zap.Error(err))
		return nil, fError.New(fError.CodeNoData)
	}
	data := make([]*models.ApiPostDetail, 0, len(page.Data))
	for _, p := range page.Data {
		user, err := dao.SelectUserById(p.AuthorId)
		if err != nil {
			zap.L().Error("SelectUserById failed.", zap.Error(err))
			return nil, fError.New(fError.CodeNoData)
		}
		community, err := dao.SelectCommunityById(p.CommunityId)
		if err != nil {
			zap.L().Error("SelectCommunityById failed.", zap.Error(err))
			return nil, fError.New(fError.CodeNoData)
		}
		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			Post:       p,
			Community:  community,
		}
		data = append(data, postDetail)
	}
	resultPage = &models.Page[*models.ApiPostDetail]{
		CurrentPage: page.CurrentPage,
		PageSize:    page.PageSize,
		Total:       page.Total,
		Pages:       page.Pages,
		Data:        data,
	}
	return
}
