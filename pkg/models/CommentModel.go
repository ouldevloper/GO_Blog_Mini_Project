package models

import "gorm.io/gorm"

type CommentModel struct {
	gorm.Model
	Id      int
	PostID  int
	Comment string
}
