package routes

import (
	"blog/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var BlogRouter = func(router *gin.Engine) {
	var v1 = router.Group("v1")
	{
		var post = v1.Group("/post")
		{
			post.POST("/post", controllers.CreatePost)
			post.GET("/posts/:Page/:Limit", controllers.GetPostsByPage)
			post.GET("/post/:PostId", controllers.GetPostById)
			post.DELETE("/post/delete/:PostId", controllers.DeletePost)
			post.PUT("/post/update/:PostId", controllers.UpdatePost)
			post.GET("/post/search/:search/:Page/:Limit", controllers.SearchForPosts)
		}

		var auth = v1.Group("auth")
		{
			auth.POST("/register", controllers.Register)

		}

	}
}
