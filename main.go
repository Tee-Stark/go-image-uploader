package main

import (
	"github.com/gin-gonic/gin"
	"go-image-uploader/config"
	"go-image-uploader/routes"
)

func main() {
	r := gin.Default()
	config.LoadEnv()

	routes.SetupUserRouter(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my Go image uploader application! 🚀",
		})
	})

	r.Run(":5000")
}
