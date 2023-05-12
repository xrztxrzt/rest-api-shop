package repository

import (
	"rest-api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
	CreateRole(role models.Role) (int, error)
}

type Products interface {
	Create(userId int, product models.Product) (int, error)
	GetAll(userId int, product models.Product) ([]models.Product, error)
	GetById(userId int, productId int) (models.Product, error)
	Delete(userId, productId int) error
	Update(userId, productId int, input models.UpdateProductInput) error
}

type Repository struct {
	Products
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Products:      NewProductPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
