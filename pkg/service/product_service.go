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
