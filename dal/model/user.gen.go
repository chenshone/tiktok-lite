// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username        string    `gorm:"column:username;not null" json:"username"`
	Password        string    `gorm:"column:password;not null" json:"password"`
	FollowCount     int32     `gorm:"column:follow_count" json:"follow_count"`     // 关注数
	FollowerCount   int32     `gorm:"column:follower_count" json:"follower_count"` // 粉丝数
	Avatar          string    `gorm:"column:avatar;not null" json:"avatar"`
	CreateAt        time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt        time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP" json:"update_at"`
	BackgroundImage string    `gorm:"column:background_image" json:"background_image"`
	Signature       string    `gorm:"column:signature" json:"signature"`
	TotalFavorited  int32     `gorm:"column:total_favorited" json:"total_favorited"` // 获赞数量
	WorkCount       int32     `gorm:"column:work_count" json:"work_count"`           // 作品数量
	FavoriteCount   int32     `gorm:"column:favorite_count" json:"favorite_count"`   // 点赞数量
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
