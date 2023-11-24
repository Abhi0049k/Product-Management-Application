package controllers

import (
	"Product-Management-Application/helpers"
	"Product-Management-Application/initializers"
	"Product-Management-Application/models"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// To fetch all the products present in the db

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := initializers.DB.Find(&products).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, products)
	}
}

// To create new Product

func CreateProduct(c *gin.Context) {
	var body struct {
		UserID             uint           `json:"user_id"`
		ProductName        string         `json:"product_name"`
		ProductDescription string         `json:"product_description"`
		ProductImages      pq.StringArray `json:"product_images" gorm:"type:text[]"`
		ProductPrice       uint32         `json:"product_price"`
		CreatedAt          time.Time
		UpdatedAt          time.Time
	}
	c.BindJSON(&body)

	body.CreatedAt = time.Now()

	body.UpdatedAt = time.Now()

	// Running the image analysis

	localFilePaths, compressedFilePaths, err := helpers.DownloadImages(body.ProductImages)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	// Storing the product in the database

	product := models.Product{
		ProductName:             body.ProductName,
		ProductDescription:      body.ProductDescription,
		ProductImages:           localFilePaths,
		ProductPrice:            body.ProductPrice,
		CreatedAt:               body.CreatedAt,
		UpdatedAt:               body.UpdatedAt,
		CompressedProductImages: compressedFilePaths,
	}

	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"msg":     "Product Created",
		"product": product,
	})
}
