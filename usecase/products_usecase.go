package usecase

import (
	"products-api/model"
	"products-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetAllProducts()
}

func (pu *ProductUsecase) SaveProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.SaveProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.Id = productId

	return product, nil
}
