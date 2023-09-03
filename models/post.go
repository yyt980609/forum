package models

import "time"

// Post 帖子详情
type Post struct {
	PostId      int64     `json:"postId,string" db:"post_id"`
	AuthorId    int64     `json:"authorId" db:"author_id"`
	CommunityId int64     `json:"communityId" db:"community_id" binding:"required" zh:"社区"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required" zh:"标题"`
	Content     string    `json:"content" db:"content" binding:"required" zh:"内容"`
	CreateTime  time.Time `json:"createTime" db:"create_time"`
}

// ApiPostDetail 帖子详情，前端展示使用
type ApiPostDetail struct {
	AuthorName string `json:"authorName"`
	*Post
	*Community `json:"community"`
}
