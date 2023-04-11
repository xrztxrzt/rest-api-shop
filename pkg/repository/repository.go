package repository

import (
	"rest-api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type Products interface {
	Create(product models.Product) (int, error)
	GetAll(product models.Product) ([]models.Product, error)
	GetById(productId int) (models.Product, error)
	Delete(productId int) error
	Update(productId int, input models.UpdateProductInput) error
}

type Repository struct {
	Products Products
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Products: NewProductPostgres(db),
	}
}
