package migrations

import (
	"blog/pkg/config"
	"blog/pkg/models"
)

func Migrate() {
	db := config.GetCnx()
	db.AutoMigrate(models.Post{}, models.Comment{}, models.Like{})
}
