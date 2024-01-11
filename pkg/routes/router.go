package routes

import (
	"blog/pkg/controllers"
	"blog/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

var BlogRouter = func(router *gin.Engine) {
	var v1 = router.Group("v1")
	{
		v1.GET("/posts/:Page/:Limit", controllers.GetPostsByPage)
		var post = v1.Group("/post")
		{
			post.GET("/:PostId", controllers.GetPostById)
			post.GET("/search/:search/:Page/:Limit", controllers.SearchForPosts)

			post.Use(middlewares.AuthMiddleware())
			{
				post.POST("/create", controllers.CreatePost)
				post.DELETE("/delete/:PostId", controllers.DeletePost)
				post.PUT("/update/:PostId", controllers.UpdatePost)
			}

		}

		var auth = v1.Group("auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)

		}

	}
}
