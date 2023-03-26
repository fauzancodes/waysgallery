package models

import "time"

type User struct {
	ID         int                     `json:"id" gorm:"primary_key:auto_increment"`
	Name       string                  `json:"name" gorm:"type: varchar(255)"`
	Email      string                  `json:"email" gorm:"type: varchar(255)"`
	Password   string                  `json:"-" gorm:"type: varchar(255)"`
	Profile    ProfileUserResponse     `json:"profile"`
	// Followings []FollowingUserResponse `json:"followings"`
	// Posts      []PostUserResponse      `json:"posts"`
	// Offers     []OfferUserResponse     `json:"offers"`
	// Orders     []OrderUserResponse     `json:"orders"`
	CreatedAt  time.Time               `json:"-"`
	UpdatedAt  time.Time               `json:"-"`
}

type UserProfileResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (UserProfileResponse) TableName() string {
	return "users"
}

type UserFollowingResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (UserFollowingResponse) TableName() string {
	return "users"
}

type UserPostResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (UserPostResponse) TableName() string {
	return "users"
}

type UserOfferResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (UserOfferResponse) TableName() string {
	return "users"
}

type UserOrderResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (UserOrderResponse) TableName() string {
	return "users"
}
