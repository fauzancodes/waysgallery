package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.WaysGalleryUser, error)
	GetUser(ID int) (models.WaysGalleryUser, error)
	UpdateUser(user models.WaysGalleryUser) (models.WaysGalleryUser, error)
	DeleteUser(user models.WaysGalleryUser) (models.WaysGalleryUser, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.WaysGalleryUser, error) {
	var users []models.WaysGalleryUser
	err := r.db.Preload("Profile").Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.WaysGalleryUser, error) {
	var user models.WaysGalleryUser
	err := r.db.Preload("Profile").First(&user, ID).Error

	return user, err
}

func (r *repository) UpdateUser(user models.WaysGalleryUser) (models.WaysGalleryUser, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.WaysGalleryUser) (models.WaysGalleryUser, error) {
	err := r.db.Delete(&user).Scan(&user).Error

	return user, err
}
