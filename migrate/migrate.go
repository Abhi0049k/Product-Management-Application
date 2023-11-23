package main

import (
	"Product-Management-Application/initializers"
	"Product-Management-Application/models"
)

func init() {
	initializers.LoadEnv()
	initializers.DBConnection()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Product{})
}
