package service

import (
	"rest-api/pkg/models"
	"rest-api/pkg/repository"
)

type ShopProductService struct {
	repo repository.Products
}

func NewProductService(repo repository.Products) *ShopProductService {
	return &ShopProductService{repo: repo}
}

func (s *ShopProductService) Create(product models.Product) (int, error) {
	return s.repo.Create(product)
}

func (s *ShopProductService) GetAll(product models.Product) ([]models.Product, error) {
	return s.repo.GetAll(product)
}

func (s *ShopProductService) GetById(productId int) (models.Product, error) {
	return s.repo.GetById(productId)
}

func (s *ShopProductService) Delete(productId int) error {
	return s.repo.Delete(productId)
}

func (s *ShopProductService) Update(productId int, input models.UpdateProductInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(productId, input)
}
