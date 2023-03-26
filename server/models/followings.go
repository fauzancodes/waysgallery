package models

import "time"

type Following struct {
	ID          int                   `json:"id" gorm:"primary_key:auto_increment"`
	FollowingID int                   `json:"following_id" gorm:"int"`
	UserID      int                   `json:"-"`
	User        UserFollowingResponse `json:"user"`
	CreatedAt   time.Time             `json:"-"`
	UpdatedAt   time.Time             `json:"-"`
}

type FollowingUserResponse struct {
	ID          int `json:"id" gorm:"primary_key:auto_increment"`
	FollowingID int `json:"following_id" gorm:"int"`
	UserID      int `json:"-"`
}

func (FollowingUserResponse) TableName() string {
	return "followings"
}
