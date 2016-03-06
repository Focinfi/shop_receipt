package models

import (
	"fmt"
)

type FavorableLineItem struct {
	ProductBarCode string
	Quantity       int
	Type           string
}

func (f FavorableLineItem) String() string {
	product, ok := products[f.ProductBarCode]
	if !ok {
		panic(fmt.Sprintf("FavorableLineItem#String: has no product with BarCode %s", f.ProductBarCode))
	}
	return fmt.Sprintf("name: %s, quantity: %d", product.Name, f.Quantity)
}
