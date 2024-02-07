package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type ArtRepository interface {
	FindArts() ([]models.WaysGalleryArt, error)
	GetArt(ID int) (models.WaysGalleryArt, error)
	CreateArt(Art models.WaysGalleryArt) (models.WaysGalleryArt, error)
	DeleteArt(Art models.WaysGalleryArt) (models.WaysGalleryArt, error)
}

func RepositoryArt(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArts() ([]models.WaysGalleryArt, error) {
	var arts []models.WaysGalleryArt
	err := r.db.Preload("Profile").Order("id desc").Find(&arts).Error

	return arts, err
}

func (r *repository) GetArt(ID int) (models.WaysGalleryArt, error) {
	var art models.WaysGalleryArt
	err := r.db.Preload("Profile").First(&art, ID).Error

	return art, err
}

func (r *repository) CreateArt(art models.WaysGalleryArt) (models.WaysGalleryArt, error) {
	err := r.db.Create(&art).Error

	return art, err
}

func (r *repository) DeleteArt(art models.WaysGalleryArt) (models.WaysGalleryArt, error) {
	err := r.db.Delete(&art).Scan(&art).Error

	return art, err
}
