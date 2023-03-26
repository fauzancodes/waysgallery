package models

import "time"

type Art struct {
	ID               int                `json:"id" gorm:"primary_key:auto_increment"`
	Image            string             `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID    string             `json:"image_public_id" gorm:"type: varchar(255)"`
	ProfileID        int                `json:"-"`
	Profile          ProfileArtResponse `json:"profile"`
	CreatedAt        time.Time          `json:"-"`
	UpdatedAt        time.Time          `json:"-"`
}

type ArtProfileResponse struct {
	ID               int    `json:"id" gorm:"primary_key:auto_increment"`
	Image            string `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID    string `json:"image_public_id" gorm:"type: varchar(255)"`
	ProfileID        int    `json:"-"`
}

func (ArtProfileResponse) TableName() string {
	return "arts"
}
