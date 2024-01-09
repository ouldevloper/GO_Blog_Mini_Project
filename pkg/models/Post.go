package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	Id          int        ` json:"id"`
	Title       string     ` json:"title"`
	Description string     ` json:"description"`
	Body        string     ` json:"body"`
	Comments    []*Comment ` json:"comments"`
	Likes       []*Like    ` jsonn:"likes"`
	User        *User      `json:"user"`
}
