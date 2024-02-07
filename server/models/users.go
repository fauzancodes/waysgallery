package models

import "time"

type WaysGalleryUser struct {
	ID       int                            `json:"id" gorm:"primary_key:auto_increment"`
	Name     string                         `json:"name" gorm:"type: varchar(255)"`
	Email    string                         `json:"email" gorm:"type: varchar(255)"`
	Password string                         `json:"-" gorm:"type: varchar(255)"`
	Profile  WaysGalleryProfileUserResponse `json:"profile"`
	// Followings []FollowingUserResponse `json:"followings"`
	// Posts      []PostUserResponse      `json:"posts"`
	// Offers     []OfferUserResponse     `json:"offers"`
	// Orders     []OrderUserResponse     `json:"orders"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type WaysGalleryUserProfileResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (WaysGalleryUserProfileResponse) TableName() string {
	return "ways_gallery_users"
}

type WaysGalleryUserFollowingResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (WaysGalleryUserFollowingResponse) TableName() string {
	return "users"
}

type WaysGalleryUserPostResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (WaysGalleryUserPostResponse) TableName() string {
	return "ways_gallery_users"
}

type WaysGalleryUserOfferResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (WaysGalleryUserOfferResponse) TableName() string {
	return "ways_gallery_users"
}

type WaysGalleryUserOrderResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (WaysGalleryUserOrderResponse) TableName() string {
	return "ways_gallery_users"
}
