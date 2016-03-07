package models

type Product struct {
	BarCode string
	Name    string
	Price   float64
}

var products = map[string]Product{}
