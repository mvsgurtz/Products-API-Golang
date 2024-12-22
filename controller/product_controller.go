package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"products-api/model"
	"products-api/usecase"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetAllProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			1, "Arroz", 20.0,
		},
	}
	ctx.JSON(http.StatusOK, products)

}
