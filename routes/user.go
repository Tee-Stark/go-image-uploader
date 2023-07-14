package routes

import (
	"github.com/gin-gonic/gin"
	"go-image-uploader/controllers"
	"go-image-uploader/middlewares"
)

func SetupUserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/:id", controllers.GetUser)
		user.POST("/:id/uploadImage", middlewares.FileUploadMiddleware(), controllers.UploadImage)
	}
}
