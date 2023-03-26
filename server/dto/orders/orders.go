package ordersdto

import "time"

type OrderRequest struct {
	Title       string    `json:"title" form:"title" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`
	StartDate   time.Time `json:"start_date" form:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" form:"end_date" validate:"required"`
	Price       int       `json:"price" form:"price" validate:"required"`
}

type OrderResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Price       int       `json:"price"`
	VendorID    int       `json:"vendor_id"`
	ClientID    int       `json:"client_id"`
	Status      string    `json:"status"`
}

type OrderStatusRequest struct {
	Status string `json:"status"`
}
