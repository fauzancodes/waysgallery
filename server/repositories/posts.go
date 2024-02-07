package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindPosts() ([]models.WaysGalleryPost, error)
	GetPost(ID int) (models.WaysGalleryPost, error)
	CreatePost(post models.WaysGalleryPost) (models.WaysGalleryPost, error)
	DeletePost(post models.WaysGalleryPost) (models.WaysGalleryPost, error)
	UpdatePost(post models.WaysGalleryPost) (models.WaysGalleryPost, error)
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPosts() ([]models.WaysGalleryPost, error) {
	var posts []models.WaysGalleryPost
	err := r.db.Preload("User").Order("id desc").Find(&posts).Error

	return posts, err
}

func (r *repository) GetPost(ID int) (models.WaysGalleryPost, error) {
	var post models.WaysGalleryPost
	err := r.db.Preload("User").First(&post, ID).Error

	return post, err
}

func (r *repository) CreatePost(post models.WaysGalleryPost) (models.WaysGalleryPost, error) {
	err := r.db.Create(&post).Error

	return post, err
}

func (r *repository) DeletePost(post models.WaysGalleryPost) (models.WaysGalleryPost, error) {
	err := r.db.Delete(&post).Scan(&post).Error

	return post, err
}

func (r *repository) UpdatePost(post models.WaysGalleryPost) (models.WaysGalleryPost, error) {
	err := r.db.Save(&post).Error

	return post, err
}
