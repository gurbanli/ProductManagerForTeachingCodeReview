package models

type Product struct {
	Model
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Price   float64 `json:"price"`
	InStock uint    `json:"in_stock"`
	UserId  uint    `json:"-"`
}
