package main

import (
	"github.com/gin-gonic/gin"
	"go-image-uploader/config"
	"go-image-uploader/models"
	"go-image-uploader/routes"
)

func main() {
	// init in-memory database
	models.NewUserDB()
	// init router
	r := gin.Default()
	config.LoadEnv()

	routes.SetupUserRouter(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my Go image uploader application! 🚀",
		})
	})

	err := r.Run(":5000")
	if err != nil {
		panic(err)
	}
}
