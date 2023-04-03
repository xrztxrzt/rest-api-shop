package models

type Cart struct {
	ID       int           `json:"id"`
	UserID   int           `json:"userId"`
	Date     string        `json:"date"`
	Products []ProductInfo `json:"products"`
	V        int           `json:"__v"`
}

type ProductInfo struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}
