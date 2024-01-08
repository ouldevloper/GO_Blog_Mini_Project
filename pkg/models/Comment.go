package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      int    `json:"id"`
	PostID  int    `json:"post_id" validate:"required"`
	Comment string `json:"comment" validate:"required"`
}
