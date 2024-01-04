package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("post", func(cxt *gin.Context) {
			cxt.JSON(200, ":hello gin ....")
		})
	}
	fmt.Println("Server run on port 3000")
	router.Run(":3000")
}
