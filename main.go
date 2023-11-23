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

	// Routes for Users
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)

	// Routes for Products
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)

	r.Run()
}
