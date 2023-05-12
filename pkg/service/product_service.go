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

func (s *ShopProductService) Create(userId int, product models.Product) (int, error) {
	return s.repo.Create(userId, product)
}

func (s *ShopProductService) GetAll(userId int, product models.Product) ([]models.Product, error) {
	return s.repo.GetAll(userId, product)
}

func (s *ShopProductService) GetById(userId, productId int) (models.Product, error) {
	return s.repo.GetById(userId, productId)
}

func (s *ShopProductService) Delete(userId, productId int) error {
	return s.repo.Delete(userId, productId)
}

func (s *ShopProductService) Update(userId, productId int, input models.UpdateProductInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, productId, input)
}
