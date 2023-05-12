package service

import (
	"rest-api/pkg/models"
	"rest-api/pkg/repository"
)

type Authorization interface {
	CreateRole(role models.Role) (int, error)
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string, roleID int) (string, error)
	ParseToken(token string) (int, string, error)
}

type ProductList interface {
	Create(userId int, product models.Product) (int, error)
	GetAll(userId int, product models.Product) ([]models.Product, error)
	GetById(userId int, productId int) (models.Product, error)
	Delete(userId, productId int) error
	Update(userId, productId int, input models.UpdateProductInput) error
}

type Service struct {
	Authorization
	ProductList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{

		Authorization: NewAuthService(repos.Authorization),
		ProductList:   NewProductService(repos.Products),
	}
}
