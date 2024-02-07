package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.WaysGalleryUser) (models.WaysGalleryUser, error)
	Login(email string) (models.WaysGalleryUser, error)
	CheckAuth(ID int) (models.WaysGalleryUser, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.WaysGalleryUser) (models.WaysGalleryUser, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(email string) (models.WaysGalleryUser, error) {
	var user models.WaysGalleryUser
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) CheckAuth(ID int) (models.WaysGalleryUser, error) {
	var user models.WaysGalleryUser
	err := r.db.Preload("Profile").Preload("Following").Preload("Post").Preload("Offer").Preload("Order").First(&user, ID).Error

	return user, err
}
