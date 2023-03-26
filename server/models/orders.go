package models

import "time"

type Order struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title       string               `json:"title" gorm:"type: varchar(255)"`
	Description string               `json:"description" gorm:"type: text"`
	StartDate   time.Time            `json:"start_date"`
	EndDate     time.Time            `json:"end_date"`
	Price       int                  `json:"price" gorm:"type: int"`
	VendorID    int                  `json:"vendor_id" gorm:"type: int"`
	ClientID    int                  `json:"client_id" gorm:"type: int"`
	Status      string               `json:"status" gorm:"type: varchar(255)"`
	Project     ProjectOrderResponse `json:"project"`
	UserID      int                  `json:"-"`
	User        UserProfileResponse  `json:"user"`
	CreatedAt   time.Time            `json:"-"`
	UpdatedAt   time.Time            `json:"-"`
}

type OrderUserResponse struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title       string               `json:"title" gorm:"type: varchar(255)"`
	Description string               `json:"description" gorm:"type: text"`
	StartDate   time.Time            `json:"start_date"`
	EndDate     time.Time            `json:"end_date"`
	Price       int                  `json:"price" gorm:"type: int"`
	VendorID    int                  `json:"vendor_id" gorm:"type: int"`
	ClientID    int                  `json:"client_id" gorm:"type: int"`
	Status      string               `json:"status" gorm:"type: varchar(255)"`
	Project     ProjectOrderResponse `json:"project" gorm:"foreignKey:OrderID"`
	UserID      int                  `json:"-"`
}

func (OrderUserResponse) TableName() string {
	return "orders"
}

type OrderProjectResponse struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Title       string    `json:"title" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type: varchar(255)"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	VendorID    int       `json:"vendor_id"`
	ClientID    int       `json:"client_id" gorm:"type: int"`
	Status      string    `json:"status" gorm:"type: varchar(255)"`
}

func (OrderProjectResponse) TableName() string {
	return "orders"
}
