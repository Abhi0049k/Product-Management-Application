package controllers

import (
	"Product-Management-Application/initializers"
	"Product-Management-Application/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// To fetch all the users present in the db

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := initializers.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

// To create a new user

func CreateUser(c *gin.Context) {
	var body struct {
		Name      string
		Mobile    string
		Latitude  float32
		Longitude float32
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	c.BindJSON(&body)
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
		c.Status(400)
		return
	}

	c.JSON(200, user)
}
