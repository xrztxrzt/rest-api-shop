package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type ProductList interface {
}

type CartList interface {
}

type Repository struct {
	Authorization
	ProductList
	CartList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
