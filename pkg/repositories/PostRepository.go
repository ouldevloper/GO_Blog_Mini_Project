package repositories

import (
	"blog/pkg/config"
	"blog/pkg/models"
)

type PostRepository struct{}

func (repos *PostRepository) Create(post *models.Post) {
	db := config.GetCnx()
	db.Create(&post)
}

func (repos *PostRepository) GetById(id int) models.Post {
	var post models.Post
	config.GetCnx().
		Model(models.Post{}).
		Where("ID=?", id).
		Find(&post)
	return post
}

func (repos *PostRepository) GetByPage(page int, offset int) []models.Post {
	var posts []models.Post
	config.GetCnx().
		Model(models.Post{}).
		Preload("Comments", "Likes").
		Limit(offset).
		Offset((page - 1) * offset).
		Find(&posts)
	return posts
}

func (repos *PostRepository) Delete(id int) {
	config.GetCnx().Delete(&models.Post{}, id)
}

func (repos *PostRepository) Update(post models.Post) {
	config.GetCnx().Save(&post)
}

func (repos *PostRepository) Search(search string, page int, offset int) []models.Post {
	var posts []models.Post
	config.GetCnx().
		Model(models.Post{}).
		Preload("Comments", "Likes").
		Where("title like ? ", search+"%").
		Or("id = ?", search).
		Or("description like ?", search+"%").
		Limit(offset).
		Offset((page - 1) * offset).
		Find(&posts)
	return posts
}
