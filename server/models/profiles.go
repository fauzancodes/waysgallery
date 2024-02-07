package models

import "time"

type WaysGalleryProfile struct {
	ID            int                             `json:"id" gorm:"primary_key:auto_increment"`
	Name          string                          `json:"name" gorm:"type: varchar(255)"`
	Greeting      string                          `json:"greeting" gorm:"type: varchar(255)"`
	Image         string                          `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID string                          `json:"image_public_id" gorm:"type: varchar(255)"`
	Arts          []WaysGalleryArtProfileResponse `json:"arts" gorm:"foreignkey:ProfileID"`
	UserID        int                             `json:"-"`
	User          WaysGalleryUserProfileResponse  `json:"user"`
	CreatedAt     time.Time                       `json:"-"`
	UpdatedAt     time.Time                       `json:"-"`
}

type WaysGalleryProfileUserResponse struct {
	ID            int                             `json:"id" gorm:"primary_key:auto_increment"`
	Name          string                          `json:"name" gorm:"type: varchar(255)"`
	Greeting      string                          `json:"greeting" gorm:"type: varchar(255)"`
	Image         string                          `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID string                          `json:"image_public_id" gorm:"type: varchar(255)"`
	Arts          []WaysGalleryArtProfileResponse `json:"arts" gorm:"foreignKey:ProfileID"`
	UserID        int                             `json:"-"`
}

func (WaysGalleryProfileUserResponse) TableName() string {
	return "ways_gallery_profiles"
}

type WaysGalleryProfileArtResponse struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Name          string `json:"name" gorm:"type: varchar(255)"`
	Greeting      string `json:"greeting" gorm:"type: varchar(255)"`
	Image         string `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID string `json:"image_public_id" gorm:"type: varchar(255)"`
}

func (WaysGalleryProfileArtResponse) TableName() string {
	return "ways_gallery_profiles"
}
