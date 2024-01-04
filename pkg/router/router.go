package router

import (
	"blog/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var BlogRouter = func(router *gin.Engine) {
	var v1 = router.Group("v1")
	{
		v1.POST("/post", controllers.CreatePost)
		v1.GET("/posts/:{Page}", controllers.GetPostsByPage)
		v1.GET("/post/{PostId}", controllers.GetPostById)
		v1.DELETE("/post/delete/{PostId}", controllers.DeletePost)
		v1.PUT("/post/update/{PostId}", controllers.UpdatePost)
		v1.GET("/post/search", controllers.SearchForPosts)
	}
}
