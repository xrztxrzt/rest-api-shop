package repository

import (
	"rest-api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Products interface {
	Create(product models.Product) (int, error)
}

type CartList interface {
}

type Repository struct {
	Authorization
	Products
	CartList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Products: NewProductPostgres(db),
	}
}
