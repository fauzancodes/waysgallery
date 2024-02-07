package models

import "time"

type WaysGalleryArt struct {
	ID            int                           `json:"id" gorm:"primary_key:auto_increment"`
	Image         string                        `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID string                        `json:"image_public_id" gorm:"type: varchar(255)"`
	ProfileID     int                           `json:"-"`
	Profile       WaysGalleryProfileArtResponse `json:"profile"`
	CreatedAt     time.Time                     `json:"-"`
	UpdatedAt     time.Time                     `json:"-"`
}

type WaysGalleryArtProfileResponse struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Image         string `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID string `json:"image_public_id" gorm:"type: varchar(255)"`
	ProfileID     int    `json:"-" gorm:"index"`
}

func (WaysGalleryArtProfileResponse) TableName() string {
	return "ways_gallery_arts"
}
