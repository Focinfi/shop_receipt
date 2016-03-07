package models

import (
	"github.com/Focinfi/shop_receipt/libs"
)

type LineItem struct {
	ProductBarCode string
	Quantity       int
	Product        Product
}

// NewLineItem allocates and returns a new LineItem with the given barCode and quantity.
// It will panic when the product with the given barCode does not exist.
func NewLineItem(barCode string, quantity int) *LineItem {
	lineItem := &LineItem{ProductBarCode: barCode, Quantity: quantity}
	if product, ok := products[barCode]; !ok {
		panic("NewLineItem: " + "has not product with bar code: " + barCode)
	} else {
		lineItem.Product = product
	}
	return lineItem
}

// CostSaving calculates and returns the cost saving with all related promotions
func (l LineItem) CostSaving() float64 {
	originalSubtotal := l.Product.Price * float64(l.Quantity)
	promotions := FindAllPromotions(l.Product.BarCode)
	costSaving := float64(0)
	for _, promotion := range promotions {
		costSaving += originalSubtotal - promotion.SubTotal(l.Product.Price, l.Quantity)
	}
	return libs.Round(costSaving, 2)
}

// Subtotal calculates and returns the subtotal
func (l LineItem) Subtotal() float64 {
	originalSubtotal := l.Product.Price * float64(l.Quantity)
	return libs.Round(originalSubtotal-l.CostSaving(), 2)
}
