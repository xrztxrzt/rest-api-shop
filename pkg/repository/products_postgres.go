package repository

import (
	"fmt"
	"rest-api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product models.Product) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int
	sqlRaw := "INSERT INTO %s (title, price, description, category, image, rating) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	sqlQuery := fmt.Sprintf(sqlRaw, ProductsTable)

	row := tx.QueryRow(sqlQuery,
		product.Title, product.Price, product.Description, product.Category, product.Image, product.Rating)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}

	return id, tx.Commit()
}
