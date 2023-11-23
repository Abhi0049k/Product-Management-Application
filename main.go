package main

import (
	"Product-Management-Application/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.DBConnection()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})
	r.Run()
}
