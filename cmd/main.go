package main

import (
	config "blog/pkg/Config"
	"blog/pkg/models"
	"blog/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("test", func(cxt *gin.Context) {
			var post *models.Post
			config.Connect()
			db := config.GetDB()
			db.Preload("like").Preload("comment").Where("id=?", 1).Find(&post)
			//utils.ParseBody(cxt.Request, &post)
			utils.Success(cxt, post, 200)
		})
	}
	fmt.Println("Server run on port 3000")
	router.Run(":3000")
}
