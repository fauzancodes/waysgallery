package models

import "time"

type WaysGalleryFollowing struct {
	ID          int                              `json:"id" gorm:"primary_key:auto_increment"`
	FollowingID int                              `json:"following_id" gorm:"int"`
	UserID      int                              `json:"-"`
	User        WaysGalleryUserFollowingResponse `json:"user"`
	CreatedAt   time.Time                        `json:"-"`
	UpdatedAt   time.Time                        `json:"-"`
}

type WaysGalleryFollowingUserResponse struct {
	ID          int `json:"id" gorm:"primary_key:auto_increment"`
	FollowingID int `json:"following_id" gorm:"int"`
	UserID      int `json:"-"`
}

func (WaysGalleryFollowingUserResponse) TableName() string {
	return "ways_gallery_followings"
}
