package models

import (
	"fmt"
	"github.com/Focinfi/shop_receipt/libs"
)

type LineItem struct {
	ProductBarCode string
	Quantity       int
	product        Product
}

func (l LineItem) CostSaving() float64 {
	originalSubtotal := l.product.Price * float64(l.Quantity)
	promotions := FindAllPromotions(l.product.BarCode)
	costSaving := float64(0)
	for _, promotion := range promotions {
		costSaving += originalSubtotal - promotion.SubTotal(l.product.Price, l.Quantity)
	}
	return libs.Round(costSaving, 2)
}

func (l LineItem) Subtotal() float64 {
	originalSubtotal := l.product.Price * float64(l.Quantity)
	return libs.Round(originalSubtotal-l.CostSaving(), 2)
}

func NewLineItem(barCode string, quantity int) *LineItem {
	lineItem := &LineItem{ProductBarCode: barCode, Quantity: quantity}
	if product, ok := products[barCode]; ok {
		lineItem.product = product
	}
	return lineItem
}

func (l *LineItem) String() string {
	output := fmt.Sprintf("name: %s, quantity: %d, unit price: %.2f(%s), subtotal: %.2f(%s)",
		l.product.Name, l.Quantity, l.product.Price, l.product.Currency, l.Subtotal(), l.product.Currency,
	)

	if l.CostSaving() > 0 {
		output = fmt.Sprintf("%s, cost saving: %.2f(%s)", output, l.CostSaving(), l.product.Currency)
	}

	return output
}
