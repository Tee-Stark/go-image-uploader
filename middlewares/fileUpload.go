package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileUploadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad request",
			})
			return
		}
		defer file.Close()

		openedFile, err := header.Open()
		if err != nil {
			fmt.Errorf("%v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "An internal error occured",
			})
		}

		// pass the file and its name to the controller
		c.Set("filePath", header.Filename)
		c.Set("file", openedFile)

		// continue to controller
		c.Next()
	}
}
