package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindProjects() ([]models.Project, error)
	GetProject(ID int) (models.Project, error)
	CreateProject(project models.Project) (models.Project, error)
}

func RepositoryProject(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProjects() ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Preload("Order").Find(&projects).Error

	return projects, err
}

func (r *repository) GetProject(ID int) (models.Project, error) {
	var project models.Project
	err := r.db.Preload("Order").First(&project, ID).Error

	return project, err
}

func (r *repository) CreateProject(project models.Project) (models.Project, error) {
	err := r.db.Create(&project).Error

	return project, err
}
