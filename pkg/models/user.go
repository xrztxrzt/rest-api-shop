package models

type User struct {
	Address  Address `json:"address"`
	City     string  `json:"city"`
	Street   string  `json:"street"`
	Number   int     `json:"number"`
	Zipcode  string  `json:"zipcode"`
	ID       int     `json:"id"`
	Email    string  `json:"email"`
	UserName string  `json:"username"`
	Password string  `json:"password"`
	Name     Name    `json:"name"`
	Phone    string  `json:"phone"`
	V        int     `json:"__v"`
}

type Address struct {
	Geolocation Geolocation `json:"geolocation"`
}

type Geolocation struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Name struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
