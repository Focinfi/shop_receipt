package models

type Product struct {
	BarCode string
	Name    string
	Price   float64
	Unit    string
}

var products = map[string]Product{}
