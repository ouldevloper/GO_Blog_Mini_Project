package migrations

import (
	"blog/pkg/config"
	"blog/pkg/models"
)

func Migrate() {
	db := config.GetCnx()
	// db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Claims{})
	// db.AutoMigrate(models.Post{})
	// db.AutoMigrate(models.Comment{})
	// db.AutoMigrate(models.Like{})
}
