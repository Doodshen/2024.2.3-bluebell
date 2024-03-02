package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"` //需要与数据库对应
}

type CommunityDetialList struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
