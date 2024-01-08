package controllers

import (
	"blog/pkg/models"
	"blog/pkg/utils"

	"github.com/gin-gonic/gin"
)

// var Post &models.PostModel

func CreatePost(c *gin.Context) {
	// title := c.Param("title")
	// body := c.Param("title")
	// description := c.Param("title")
	var post *models.PostModel
	utils.ParseBody(c.Request, &post)
	c.JSON(200, post)
}

func GetPostsByPage(c *gin.Context) {

}

func GetPostById(c *gin.Context) {

}

func DeletePost(c *gin.Context) {

}

func UpdatePost(c *gin.Context) {

}

func SearchForPosts(c *gin.Context) {

}
