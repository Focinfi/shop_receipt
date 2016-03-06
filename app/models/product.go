package models

type Currency string

const (
	RMB Currency = "元"
	USD Currency = "$"
)

type Product struct {
	BarCode  string
	Name     string
	Price    float64
	Currency Currency
}

var products = map[string]Product{}
