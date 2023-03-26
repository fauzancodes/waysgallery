package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type ArtRepository interface {
	FindArts() ([]models.Art, error)
	GetArt(ID int) (models.Art, error)
	CreateArt(Art models.Art) (models.Art, error)
	DeleteArt(Art models.Art) (models.Art, error)
}

func RepositoryArt(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArts() ([]models.Art, error) {
	var arts []models.Art
	err := r.db.Preload("Profile").Order("id desc").Find(&arts).Error

	return arts, err
}

func (r *repository) GetArt(ID int) (models.Art, error) {
	var art models.Art
	err := r.db.Preload("Profile").First(&art, ID).Error

	return art, err
}

func (r *repository) CreateArt(art models.Art) (models.Art, error) {
	err := r.db.Create(&art).Error

	return art, err
}

func (r *repository) DeleteArt(art models.Art) (models.Art, error) {
	err := r.db.Delete(&art).Scan(&art).Error

	return art, err
}
