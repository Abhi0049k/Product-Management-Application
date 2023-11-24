package controllers

import (
	"Product-Management-Application/initializers"
	"Product-Management-Application/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// To fetch all the users present in the db

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	if err := initializers.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(http.StatusOK, users)
}

// To create a new user

func (uc *UserController) CreateUser(c *gin.Context) {
	var body models.User
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(500)
		return
	}

	body.CreatedAt = time.Now()
	body.UpdatedAt = time.Now()

	user := models.User{
		Name:      body.Name,
		Mobile:    body.Mobile,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(http.StatusOK, user)
}
