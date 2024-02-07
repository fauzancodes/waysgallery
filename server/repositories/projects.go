package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindProjects() ([]models.WaysGalleryProject, error)
	GetProject(ID int) (models.WaysGalleryProject, error)
	CreateProject(project models.WaysGalleryProject) (models.WaysGalleryProject, error)
}

func RepositoryProject(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProjects() ([]models.WaysGalleryProject, error) {
	var projects []models.WaysGalleryProject
	err := r.db.Preload("Order").Find(&projects).Error

	return projects, err
}

func (r *repository) GetProject(ID int) (models.WaysGalleryProject, error) {
	var project models.WaysGalleryProject
	err := r.db.Preload("Order").First(&project, ID).Error

	return project, err
}

func (r *repository) CreateProject(project models.WaysGalleryProject) (models.WaysGalleryProject, error) {
	err := r.db.Create(&project).Error

	return project, err
}
