package models

type Community struct {
	Id            int64  `json:"id" db:"id"`
	CommunityId   string `json:"communityId" db:"community_id"`
	CommunityName string `json:"communityName" db:"community_name"`
	Introduction  string `json:"introduction" db:"introduction"`
}
