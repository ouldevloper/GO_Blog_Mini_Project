package migrations

import (
	config "blog/pkg/Config"
	"blog/pkg/models"
)

go func Migrate() {
	db := config.GetCnx()
	db.AutoMigrate(models.Post{}, models.Comment{}, models.Like{})
}
