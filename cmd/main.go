package main

import (
	"blog/pkg/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.BlogRouter(router)
	fmt.Println("Server run on port 3000")
	router.Run(":3000")
}
