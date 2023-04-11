package service

import (
	"rest-api/pkg/models"
	"rest-api/pkg/repository"
)

type ProductList interface {
	Create(product models.Product) (int, error)
	GetAll(product models.Product) ([]models.Product, error)
	GetById(productId int) (models.Product, error)
	Delete(productId int) error
	Update(productId int, input models.UpdateProductInput) error
}

type Service struct {
	ProductList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{

		ProductList: NewProductService(repos.Products),
	}
}
