package main

import (
	"github.com/gin-gonic/gin"
	"products-api/controller"
	"products-api/db"
	"products-api/repository"
	"products-api/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Camada UseCase
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)

	//Camada controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetAllProducts)
	server.GET("/products/:productId", ProductController.GetProductById)
	server.POST("/product", ProductController.SaveProduct)

	server.Run(":8080")
}
