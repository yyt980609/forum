package models

import "time"

type Community struct {
	Id            int64     `json:"id" db:"id"`
	CommunityId   string    `json:"communityId" db:"community_id"`
	CommunityName string    `json:"communityName" db:"community_name"`
	Introduction  string    `json:"introduction" db:"introduction"`
	CreateTime    time.Time `json:"createTime" db:"create_time"`
	UpdateTime    time.Time `json:"updateTime" db:"update_time"`
}
