package main

import (
	"github.com/gin-gonic/gin"
	"products-api/controller"
)

func main() {

	server := gin.Default()

	productController := controller.NewProductController()

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetAllProducts)

	server.Run(":8080")
}
