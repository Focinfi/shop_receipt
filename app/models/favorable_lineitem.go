package models

import (
	"fmt"
)

type FavorableLineItem struct {
	ProductBarCode string
	Quantity       int
	Type           string
}

func (f FavorableLineItem) Product() Product {
	if product, ok := products[f.ProductBarCode]; !ok {
		panic(fmt.Sprintf("FavorableLineItem#String: has no product with BarCode %s", f.ProductBarCode))
	} else {
		return product
	}
}
