package controllers

import (
	"Product-Management-Application/initializers"
	"Product-Management-Application/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := initializers.DB.Find(&products).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, products)
	}
}

func CreateProduct(c *gin.Context) {
	var body struct {
		UserID             uint           `json:"user_id"`
		ProductName        string         `json:"product_name"`
		ProductDescription string         `json:"product_description"`
		ProductImages      pq.StringArray `json:"product_images" gorm:"type:text[]"`
		ProductPrice       float64        `json:"product_price"`
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	c.BindJSON(&body)

	body.CreatedAt = time.Now()

	body.UpdatedAt = time.Now()

	product := models.Product{ProductName: body.ProductName, ProductDescription: body.ProductDescription, ProductImages: body.ProductImages, CreatedAt: body.CreatedAt, UpdatedAt: body.UpdatedAt, CompressedProductImages: body.ProductImages}

	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"msg": "Product Created",
	})
}
