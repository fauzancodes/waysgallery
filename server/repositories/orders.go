package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrders() ([]models.Order, error)
	GetOrder(ID int) (models.Order, error)
	CreateOrder(Order models.Order) (models.Order, error)
	UpdateOrderStatus(order models.Order) (models.Order, error)
	UpdateOrder(status string, orderId int) (models.Order, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("User").Preload("Project").Find(&orders).Error

	return orders, err
}

func (r *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("User").Preload("Project").First(&order, ID).Error

	return order, err
}

func (r *repository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error

	return order, err
}

func (r *repository) UpdateOrderStatus(order models.Order) (models.Order, error) {
	err := r.db.Save(&order).Error

	return order, err
}

func (r *repository) UpdateOrder(status string, orderId int) (models.Order, error) {
  var order models.Order
  r.db.Preload("User").Preload("Project").First(&order, orderId)

	var err error
  if status != order.Status && status == "waiting" {
		order.Status = status
		err = r.db.Save(&order).Error
  }
  return order, err
}