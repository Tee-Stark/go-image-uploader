package controllers

import (
	"github.com/gin-gonic/gin"
	"go-image-uploader/models"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user.CreateUser()
	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	userId, _ := strconv.Atoi(id)
	user = user.GetUser(userId)
	c.JSON(http.StatusOK, user)
}
