package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type FollowingRepository interface {
	FindFollowings() ([]models.Following, error)
	GetFollowing(ID int) (models.Following, error)
	CreateFollowing(Following models.Following) (models.Following, error)
	DeleteFollowing(Following models.Following) (models.Following, error)
}

func RepositoryFollowing(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFollowings() ([]models.Following, error) {
	var followings []models.Following
	err := r.db.Preload("User").Find(&followings).Error

	return followings, err
}

func (r *repository) GetFollowing(ID int) (models.Following, error) {
	var following models.Following
	err := r.db.Preload("User").First(&following, ID).Error

	return following, err
}

func (r *repository) CreateFollowing(following models.Following) (models.Following, error) {
	err := r.db.Create(&following).Error

	return following, err
}

func (r *repository) DeleteFollowing(following models.Following) (models.Following, error) {
	err := r.db.Delete(&following).Scan(&following).Error

	return following, err
}
