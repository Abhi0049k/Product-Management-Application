package controllers

import (
	"Product-Management-Application/helpers"
	"Product-Management-Application/initializers"
	"Product-Management-Application/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	var localFilePaths []string

	var compressedFilePaths []string

	for _, imageURL := range body.ProductImages {
		fileName := fmt.Sprintf("image_%d%s", time.Now().UnixNano(), filepath.Ext(imageURL))
		destination := filepath.Join("./images", fileName)
		destination = filepath.ToSlash(destination)
		localFilePaths = append(localFilePaths, destination)
		err := helpers.DownloadImage(imageURL, destination)
		if err != nil {
			log.Fatal("Error:", err.Error())
		}

		compressedFileName := fmt.Sprintf("compressed_%d%s", time.Now().UnixNano(), filepath.Ext(imageURL))
		compressedDestination := filepath.Join("./compressed-images", compressedFileName)
		compressedDestination = filepath.ToSlash(compressedDestination)
		err = os.MkdirAll(filepath.Dir(compressedDestination), os.ModePerm)

		if err != nil {
			log.Fatal("Error:", err.Error())
		}

		err = helpers.CompressImage(destination, compressedDestination, 100, 100)
		if err != nil {
			log.Fatal("Error:", err.Error())
		}

		compressedFilePaths = append(compressedFilePaths, compressedDestination)
	}

	product := models.Product{ProductName: body.ProductName, ProductDescription: body.ProductDescription, ProductImages: localFilePaths, CreatedAt: body.CreatedAt, UpdatedAt: body.UpdatedAt, CompressedProductImages: compressedFilePaths}

	// result := initializers.DB.Create(&product)

	// if result.Error != nil {
	// 	c.Status(400)
	// 	return
	// }
	c.JSON(200, gin.H{
		"msg":     "Product Created",
		"product": product,
	})
}
