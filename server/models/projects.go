package models

import "time"

type Project struct {
	ID             int                  `json:"id" gorm:"primary_key:auto_increment"`
	Description    string               `json:"description" gorm:"type: text"`
	Image1         string               `json:"image_1" gorm:"type: varchar(255)"`
	ImagePublicID1 string               `json:"image_public_id_1" gorm:"type: varchar(255)"`
	Image2         string               `json:"image_2" gorm:"type: varchar(255)"`
	ImagePublicID2 string               `json:"image_public_id_2" gorm:"type: varchar(255)"`
	Image3         string               `json:"image_3" gorm:"type: varchar(255)"`
	ImagePublicID3 string               `json:"image_public_id_3" gorm:"type: varchar(255)"`
	Image4         string               `json:"image_4" gorm:"type: varchar(255)"`
	ImagePublicID4 string               `json:"image_public_id_4" gorm:"type: varchar(255)"`
	Image5         string               `json:"image_5" gorm:"type: varchar(255)"`
	ImagePublicID5 string               `json:"image_public_id_5" gorm:"type: varchar(255)"`
	OrderID        int                  `json:"-"`
	Order          OrderProjectResponse `json:"order"`
	CreatedAt      time.Time            `json:"-"`
	UpdatedAt      time.Time            `json:"-"`
}

type ProjectOrderResponse struct {
	ID             int    `json:"id" gorm:"primary_key:auto_increment"`
	Description    string `json:"description" gorm:"type: text"`
	Image1         string `json:"image_1" gorm:"type: varchar(255)"`
	ImagePublicID1 string `json:"image_public_id_1" gorm:"type: varchar(255)"`
	Image2         string `json:"image_2" gorm:"type: varchar(255)"`
	ImagePublicID2 string `json:"image_public_id_2" gorm:"type: varchar(255)"`
	Image3         string `json:"image_3" gorm:"type: varchar(255)"`
	ImagePublicID3 string `json:"image_public_id_3" gorm:"type: varchar(255)"`
	Image4         string `json:"image_4" gorm:"type: varchar(255)"`
	ImagePublicID4 string `json:"image_public_id_4" gorm:"type: varchar(255)"`
	Image5         string `json:"image_5" gorm:"type: varchar(255)"`
	ImagePublicID5 string `json:"image_public_id_5" gorm:"type: varchar(255)"`
	OrderID        int    `json:"-"`
}

func (ProjectOrderResponse) TableName() string {
	return "projects"
}
