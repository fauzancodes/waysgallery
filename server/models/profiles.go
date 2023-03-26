package models

import "time"

type Profile struct {
	ID            int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name          string               `json:"name" gorm:"type: varchar(255)"`
	Greeting      string               `json:"greeting" gorm:"type: varchar(255)"`
	Image         string               `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID string               `json:"image_public_id" gorm:"type: varchar(255)"`
	Arts          []ArtProfileResponse `json:"arts"`
	UserID        int                  `json:"-"`
	User          UserProfileResponse  `json:"user"`
	CreatedAt     time.Time            `json:"-"`
	UpdatedAt     time.Time            `json:"-"`
}

type ProfileUserResponse struct {
	ID            int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name          string               `json:"name" gorm:"type: varchar(255)"`
	Greeting      string               `json:"greeting" gorm:"type: varchar(255)"`
	Image         string               `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID string               `json:"image_public_id" gorm:"type: varchar(255)"`
	Arts          []ArtProfileResponse `json:"arts" gorm:"foreignKey:ProfileID"`
	UserID        int                  `json:"-"`
}

func (ProfileUserResponse) TableName() string {
	return "profiles"
}

type ProfileArtResponse struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Name          string `json:"name" gorm:"type: varchar(255)"`
	Greeting      string `json:"greeting" gorm:"type: varchar(255)"`
	Image         string `json:"image" gorm:"type: varchar(255)"`
	ImagePublicID string `json:"image_public_id" gorm:"type: varchar(255)"`
}

func (ProfileArtResponse) TableName() string {
	return "profiles"
}
