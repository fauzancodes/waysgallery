package models

import "time"

type WaysGalleryOrder struct {
	ID          int                             `json:"id" gorm:"primary_key:auto_increment"`
	Title       string                          `json:"title" gorm:"type: varchar(255)"`
	Description string                          `json:"description" gorm:"type: text"`
	StartDate   time.Time                       `json:"start_date"`
	EndDate     time.Time                       `json:"end_date"`
	Price       int                             `json:"price" gorm:"type: int"`
	VendorID    int                             `json:"vendor_id" gorm:"type: int"`
	ClientID    int                             `json:"client_id" gorm:"type: int"`
	Status      string                          `json:"status" gorm:"type: varchar(255)"`
	Project     WaysGalleryProjectOrderResponse `json:"project" gorm:"foreignkey:OrderID"`
	UserID      int                             `json:"-"`
	User        WaysGalleryUserProfileResponse  `json:"user"`
	CreatedAt   time.Time                       `json:"-"`
	UpdatedAt   time.Time                       `json:"-"`
}

type WaysGalleryOrderUserResponse struct {
	ID          int                             `json:"id" gorm:"primary_key:auto_increment"`
	Title       string                          `json:"title" gorm:"type: varchar(255)"`
	Description string                          `json:"description" gorm:"type: text"`
	StartDate   time.Time                       `json:"start_date"`
	EndDate     time.Time                       `json:"end_date"`
	Price       int                             `json:"price" gorm:"type: int"`
	VendorID    int                             `json:"vendor_id" gorm:"type: int"`
	ClientID    int                             `json:"client_id" gorm:"type: int"`
	Status      string                          `json:"status" gorm:"type: varchar(255)"`
	Project     WaysGalleryProjectOrderResponse `json:"project" gorm:"foreignKey:OrderID"`
	UserID      int                             `json:"-"`
}

func (WaysGalleryOrderUserResponse) TableName() string {
	return "ways_gallery_orders"
}

type WaysGalleryOrderProjectResponse struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Title       string    `json:"title" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type: varchar(255)"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	VendorID    int       `json:"vendor_id"`
	ClientID    int       `json:"client_id" gorm:"type: int"`
	Status      string    `json:"status" gorm:"type: varchar(255)"`
}

func (WaysGalleryOrderProjectResponse) TableName() string {
	return "ways_gallery_orders"
}
