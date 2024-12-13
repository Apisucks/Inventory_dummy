package models

var Objects []Object

type Object struct {
	Name     string  `json:"name"`
	Quantity int     `json:"qty"`
	Price    float32 `json:"price"`
}
