package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrders() ([]models.WaysGalleryOrder, error)
	GetOrder(ID int) (models.WaysGalleryOrder, error)
	CreateOrder(Order models.WaysGalleryOrder) (models.WaysGalleryOrder, error)
	UpdateOrderStatus(order models.WaysGalleryOrder) (models.WaysGalleryOrder, error)
	UpdateOrder(status string, orderId int) (models.WaysGalleryOrder, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOrders() ([]models.WaysGalleryOrder, error) {
	var orders []models.WaysGalleryOrder
	err := r.db.Preload("User").Preload("Project").Find(&orders).Error

	return orders, err
}

func (r *repository) GetOrder(ID int) (models.WaysGalleryOrder, error) {
	var order models.WaysGalleryOrder
	err := r.db.Preload("User").Preload("Project").First(&order, ID).Error

	return order, err
}

func (r *repository) CreateOrder(order models.WaysGalleryOrder) (models.WaysGalleryOrder, error) {
	err := r.db.Create(&order).Error

	return order, err
}

func (r *repository) UpdateOrderStatus(order models.WaysGalleryOrder) (models.WaysGalleryOrder, error) {
	err := r.db.Save(&order).Error

	return order, err
}

func (r *repository) UpdateOrder(status string, orderId int) (models.WaysGalleryOrder, error) {
	var order models.WaysGalleryOrder
	r.db.Preload("User").Preload("Project").First(&order, orderId)

	var err error
	if status != order.Status && status == "waiting" {
		order.Status = status
		err = r.db.Save(&order).Error
	}
	return order, err
}
