package models

//binding:"required" -эти теги валидируют начилие данных полей в теле запроса
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `lson:"description" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Image       string  `json:"image" binding:"required"`
	Rating      Rating  `json:"rating" binding:"required"`
}

type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}
