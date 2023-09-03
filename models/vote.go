package models

// VoteData 投票信息
type VoteData struct {
	// PostId 帖子id
	PostId string `json:"postId" binding:"required" zh:"帖子ID"`
	// 投票信息，1赞成，-1反对
	Direction string `json:"direction" binding:"oneof=1 0 -1" zh:"投票结果"`
}
