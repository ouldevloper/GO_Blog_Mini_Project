package controllers

import (
	"blog/pkg/models"
	"blog/pkg/repositories"
	"blog/pkg/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	repos = repositories.PostRepository{}
)

func CreatePost(c *gin.Context) {

	// title := c.Param("title")
	// body := c.Param("title")
	// description := c.Param("title")
	// post := models.Post{
	// 	Title:       title,
	// 	Body:        body,
	// 	Description: description,
	// }
	var post models.Post
	utils.ParseBody(c.Request, &post)
	repos.Create(&post)
	utils.ParseBody(c.Request, &post)
	c.JSON(200, post)
}

func GetPostsByPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("Page"))
	limit, _ := strconv.Atoi(c.Param("Limit"))
	posts := repos.GetByPage(page, limit)
	utils.Success(c, posts, http.StatusOK)
}

func GetPostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("PostId"))
	post := repos.GetById(id)
	utils.Success(c, post, http.StatusOK)
}

func DeletePost(c *gin.Context) {
	fmt.Println("params", c.Request.Body)
	id, _ := strconv.Atoi(c.Param("PostId"))
	repos.Delete(id)
	utils.Success(c, nil, http.StatusOK)
}

func UpdatePost(c *gin.Context) {

	var post models.Post
	utils.ParseBody(c.Request, &post)
	repos.Update(post)
	utils.Success(c, nil, http.StatusOK)
}

func SearchForPosts(c *gin.Context) {
	search := c.Param("search")
	page, _ := strconv.Atoi(c.Param("Page"))
	limit, _ := strconv.Atoi(c.Param("Limit"))
	posts := repos.Search(search, page, limit)
	utils.Success(c, posts, http.StatusOK)
}
