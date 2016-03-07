package models

import (
	// "fmt"

	// "github.com/Focinfi/shop_receipt/config"
	"github.com/Focinfi/shop_receipt/libs"
)

type LineItem struct {
	ProductBarCode string
	Quantity       int
	Product        Product
}

func (l LineItem) CostSaving() float64 {
	originalSubtotal := l.Product.Price * float64(l.Quantity)
	promotions := FindAllPromotions(l.Product.BarCode)
	costSaving := float64(0)
	for _, promotion := range promotions {
		costSaving += originalSubtotal - promotion.SubTotal(l.Product.Price, l.Quantity)
	}
	return libs.Round(costSaving, 2)
}

func (l LineItem) Subtotal() float64 {
	originalSubtotal := l.Product.Price * float64(l.Quantity)
	return libs.Round(originalSubtotal-l.CostSaving(), 2)
}

func NewLineItem(barCode string, quantity int) *LineItem {
	lineItem := &LineItem{ProductBarCode: barCode, Quantity: quantity}
	if product, ok := products[barCode]; ok {
		lineItem.Product = product
	}
	return lineItem
}
