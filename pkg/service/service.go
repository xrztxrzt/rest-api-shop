package service

import (
	"rest-api/pkg/models"
	"rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type ProductList interface {
}

type CartList interface {
}
type Service struct {
	Authorization
	ProductList
	CartList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
