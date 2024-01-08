package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	// Id      int
	PostID  int    `json:"post_id" validate:"required"`
	Comment string `json:"comment" validate:"required"`
}
