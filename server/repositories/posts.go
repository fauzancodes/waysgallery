package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindPosts() ([]models.Post, error)
	GetPost(ID int) (models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	DeletePost(post models.Post) (models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPosts() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("User").Order("id desc").Find(&posts).Error

	return posts, err
}

func (r *repository) GetPost(ID int) (models.Post, error) {
	var post models.Post
	err := r.db.Preload("User").First(&post, ID).Error

	return post, err
}

func (r *repository) CreatePost(post models.Post) (models.Post, error) {
	err := r.db.Create(&post).Error

	return post, err
}

func (r *repository) DeletePost(post models.Post) (models.Post, error) {
	err := r.db.Delete(&post).Scan(&post).Error

	return post, err
}

func (r *repository) UpdatePost(post models.Post) (models.Post, error) {
	err := r.db.Save(&post).Error

	return post, err
}
