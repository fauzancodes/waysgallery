package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindProfiles() ([]models.WaysGalleryProfile, error)
	GetProfile(ID int) (models.WaysGalleryProfile, error)
	CreateProfile(profile models.WaysGalleryProfile) (models.WaysGalleryProfile, error)
	UpdateProfile(profile models.WaysGalleryProfile) (models.WaysGalleryProfile, error)
	DeleteProfile(profile models.WaysGalleryProfile) (models.WaysGalleryProfile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProfiles() ([]models.WaysGalleryProfile, error) {
	var profiles []models.WaysGalleryProfile
	err := r.db.Preload("User").Preload("Arts").Find(&profiles).Error

	return profiles, err
}

func (r *repository) GetProfile(ID int) (models.WaysGalleryProfile, error) {
	var profile models.WaysGalleryProfile
	err := r.db.Preload("User").Preload("Arts").First(&profile, ID).Error

	return profile, err
}

func (r *repository) CreateProfile(profile models.WaysGalleryProfile) (models.WaysGalleryProfile, error) {
	err := r.db.Create(&profile).Error

	return profile, err
}

func (r *repository) UpdateProfile(profile models.WaysGalleryProfile) (models.WaysGalleryProfile, error) {
	err := r.db.Save(&profile).Error

	return profile, err
}

func (r *repository) DeleteProfile(profile models.WaysGalleryProfile) (models.WaysGalleryProfile, error) {
	err := r.db.Delete(&profile).Scan(&profile).Error

	return profile, err
}
