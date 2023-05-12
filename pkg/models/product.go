package models

import "errors"

type Product struct {
	ID          int     `json:"-" db:"id"`
	Title       string  `json:"title" db:"title"`
	Price       float64 `json:"price" db:"price" binding:"required"`
	Description string  `json:"description" db:"description"`
	Category    string  `json:"category" db:"category" binding:"required"`
	Image       string  `json:"image" db:"image"`
	Done        bool    `json:"done" db:"done"`
}

type UserProduct struct {
	ID        int
	UserId    int
	ProductId int
}

type UpdateProductInput struct {
	Title       *string  `json:"title"`
	Price       *float64 `json:"price"`
	Description *string  `json:"description"`
	Category    *string  `json:"category"`
	Image       *string  `json:"image"`
	Done        *bool    `json:"done"`
}

func (i UpdateProductInput) Validate() error {
	if i.Title == nil && i.Price == nil && i.Description == nil && i.Category == nil && i.Image == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
