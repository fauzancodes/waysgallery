package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type FollowingRepository interface {
	FindFollowings() ([]models.WaysGalleryFollowing, error)
	GetFollowing(ID int) (models.WaysGalleryFollowing, error)
	CreateFollowing(Following models.WaysGalleryFollowing) (models.WaysGalleryFollowing, error)
	DeleteFollowing(Following models.WaysGalleryFollowing) (models.WaysGalleryFollowing, error)
}

func RepositoryFollowing(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFollowings() ([]models.WaysGalleryFollowing, error) {
	var followings []models.WaysGalleryFollowing
	err := r.db.Preload("User").Find(&followings).Error

	return followings, err
}

func (r *repository) GetFollowing(ID int) (models.WaysGalleryFollowing, error) {
	var following models.WaysGalleryFollowing
	err := r.db.Preload("User").First(&following, ID).Error

	return following, err
}

func (r *repository) CreateFollowing(following models.WaysGalleryFollowing) (models.WaysGalleryFollowing, error) {
	err := r.db.Create(&following).Error

	return following, err
}

func (r *repository) DeleteFollowing(following models.WaysGalleryFollowing) (models.WaysGalleryFollowing, error) {
	err := r.db.Delete(&following).Scan(&following).Error

	return following, err
}
