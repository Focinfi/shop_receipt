package models

type Product struct {
	BarCode string
	Name    string
	Price   float64
	Uint    string
}

var products = map[string]Product{}
