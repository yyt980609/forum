package service

import (
	"forum/common"
	"forum/models"
	fRedis "forum/pkg/redis"
	fError "forum/utils/forum_error"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

/*
投票算法：http://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_reddit.html
本项目使用简化版的投票分数
投一票就加432分 86400/200 -> 200张赞成票就可以给帖子在首页续天  -> 《redis实战》
*/

/*
	PostVote 为帖子投票

投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票

记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

v=1时，有两种情况

	1.之前没投过票，现在要投赞成票		--> 更新分数和投票记录		差值的绝对值：1  +432
	2.之前投过反对票，现在要改为赞成票	--> 更新分数和投票记录		差值的绝对值：2  +432*2

v=0时，有两种情况

	1.之前投过反对票，现在要取消			--> 更新分数和投票记录		差值的绝对值：1  +432
	2.之前投过赞成票，现在要取消			--> 更新分数和投票记录		差值的绝对值：1  -432

v=-1时，有两种情况

	1.之前没投过票，现在要投反对票		--> 更新分数和投票记录		差值的绝对值：1  -432
	2.之前投过赞成票，现在要改为反对票	--> 更新分数和投票记录		差值的绝对值：2  -432*2

以上分析可知，当需要加分时，之前的操作的值小于当前的值，所以可以利用这一特性，减少if else的分支，
而分值的数量，可以使用操作的差值的绝对值来计算
投票的限制：
每个帖子子发表之日起一个星期之内允许用户投票，超过一个星期就不允许投票了

	1、到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2、到期之后删除那个 KeyPostVotedZSetPrefix
*/
func Vote(userId int64, voteData *models.VoteData) error {
	postId := voteData.PostId
	direction, _ := strconv.ParseFloat(voteData.Direction, 64)
	client := fRedis.GetRedis()
	ctx := fRedis.GetContext()
	// 1.查询帖子发布时间
	postTime := client.ZScore(ctx, common.GetRedisKey(common.KeyPostTimeZSet), postId).Val()
	if float64(time.Now().Unix())-postTime > common.OneWeekInSeconds {
		return fError.New(fError.CodeCanNotVote)
	}
	// 2.查询当前用户给帖子的投票记录
	overVote := client.ZScore(ctx, common.GetRedisKey(common.KeyPostVotedPrefix+postId), strconv.FormatInt(userId, 10)).Val()
	var op float64
	if direction > overVote {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(overVote - direction)
	// 3.投票
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(ctx, common.GetRedisKey(common.KeyPostScoreZSet), op*diff*common.ScorePerVote, postId)
	// 4.记录用户的投票信息
	if direction == 0 {
		// 移除投票
		pipeline.ZRem(ctx, common.GetRedisKey(common.KeyPostVotedPrefix+postId), userId)
	} else {
		// 记录投票信息
		pipeline.ZAdd(ctx, common.GetRedisKey(common.KeyPostVotedPrefix+postId), &redis.Z{
			Score:  direction,
			Member: userId,
		})
	}
	_, err := pipeline.Exec(ctx)
	if err != nil {
		return fError.New(fError.CodeSystemError)
	}
	return err
}
