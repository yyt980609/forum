package common

// UserId 用户标识
const UserId = "userId"

// APPName 应用名
const APPName = "forum"

// TokenHeader Token Header
const TokenHeader = "Authorization"

// Bearer Toke组成
const Bearer = "Bearer"

const (
	// OneWeekInSeconds 一周的秒数
	OneWeekInSeconds = 7 * 27 * 3600

	// ScorePerVote 每一票值多少分
	ScorePerVote = 432
)

// redis key
const (
	// KeyPrefix 前缀
	KeyPrefix = "forum:"
	// KeyPostTimeZSet 按发帖时间做分数的ZSet的key
	KeyPostTimeZSet = "post:time"
	// KeyPostScoreZSet 按投票做分数的ZSet的key
	KeyPostScoreZSet = "post:score"
	// KeyPostVotedPrefix 投票
	KeyPostVotedPrefix = "post:voted:"
)

func GetRedisKey(key string) string {
	return KeyPrefix + key
}
