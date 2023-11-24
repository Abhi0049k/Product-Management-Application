package main

import (
	"Product-Management-Application/controllers"
	"Product-Management-Application/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.DBConnection()
}

func main() {
	r := gin.Default()

	userController := &controllers.UserController{}
	productController := &controllers.ProductController{}

	// Routes for Users
	r.GET("/users", userController.GetUsers)
	r.POST("/users", userController.CreateUser)

	// Routes for Products
	r.GET("/products", productController.GetProducts)
	r.POST("/products", productController.CreateProduct)

	r.Run()
}
