package repository

import (
	"fmt"
	"rest-api/pkg/models"
	"strings"

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

	sqlRaw :=
		`INSERT INTO Products (title, price, description, category, image) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	id := 0

	err = tx.QueryRow(sqlRaw,
		product.Title, product.Price, product.Description, product.Category, product.Image).Scan(&id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, nil
}

func (r *ProductPostgres) GetAll(product models.Product) ([]models.Product, error) {
	var products []models.Product

	sqlRaw := `SELECT * FROM Products`
	//sqlQuery := fmt.Sprintf(sqlRaw, ProductsTable)

	err := r.db.Select(&products, sqlRaw)

	return products, err

}

func (r *ProductPostgres) GetById(productId int) (models.Product, error) {
	var product models.Product

	sqlRaw := `SELECT id, title, price, description, category, image  FROM Products WHERE id = %d`
	sqlQuery := fmt.Sprintf(sqlRaw, productId)

	err := r.db.Get(&product, sqlQuery)
	if err != nil {
		return models.Product{}, err
	}

	return product, err
}

func (r *ProductPostgres) Delete(productId int) error {
	sqlRaw := `DELETE FROM Products WHERE id = %d`
	sqlQuery := fmt.Sprintf(sqlRaw, productId)

	_, err := r.db.Exec(sqlQuery)

	return err
}

func (r *ProductPostgres) Update(productId int, input models.UpdateProductInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	sqlRaw := `UPDATE Products set`
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Category != nil {
		setValues = append(setValues, fmt.Sprintf("category=$%d", argId))
		args = append(args, *input.Category)
		argId++
	}

	if input.Image != nil {
		setValues = append(setValues, fmt.Sprintf("image=$%d", argId))
		args = append(args, *input.Image)
		argId++
	}

	if len(setValues) == 0 {
		return nil
	}

	sqlRaw += " " + strings.Join(setValues, ", ")
	sqlRaw += " WHERE id=$" + fmt.Sprintf("%d", argId)

	args = append(args, productId)

	_, err := r.db.Exec(sqlRaw, args...)
	return err
}
