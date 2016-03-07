package models

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/Focinfi/shop_receipt/libs"
)

type Receipt struct {
	LineItems map[string]*LineItem
}

// NewReceipt allocates and returns a new Receipt with the given barCodes.
// It will treat "P100010-2" as a "P100010" with quantity 2.
func NewReceipt(barCodes []string) *Receipt {
	lineItems := map[string]*LineItem{}
	for _, barCode := range barCodes {
		quantityCode := regexp.MustCompile("-[0-9]+$").FindString(barCode)
		productBarCode := barCode
		quantity := 1
		if quantityCode != "" {
			productBarCode = strings.TrimSuffix(productBarCode, quantityCode)
			quantityCode = strings.TrimPrefix(quantityCode, "-")
			quantity, _ = strconv.Atoi(quantityCode)
		}

		if li, ok := lineItems[productBarCode]; !ok {
			lineItems[productBarCode] = NewLineItem(productBarCode, quantity)
		} else {
			li.Quantity += quantity
		}
	}

	return &Receipt{LineItems: lineItems}
}

// Total sums and returns subtotal of every LineItems
func (r Receipt) Total() float64 {
	var total float64
	for _, li := range r.LineItems {
		total += li.Subtotal()
		fmt.Println(li)
	}
	return total
}

// CostSaving sums and returns cost saving of every Lineitem
func (r Receipt) CostSaving() float64 {
	var costSaving float64
	for _, li := range r.LineItems {
		costSaving += li.CostSaving()
	}
	return costSaving
}

// FavorableLineItemMap traverses LineItems and returns all FavorableLineItem as map
func (r Receipt) FavorableLineItemMap() map[string][]FavorableLineItem {
	favorableLineItemsMap := map[string][]FavorableLineItem{}
	for _, li := range r.LineItems {
		for _, promotion := range FindAllPromotions(li.Product.BarCode) {
			if promotion.Printable {
				favorableLineItem := FavorableLineItem{
					ProductBarCode: li.Product.BarCode,
					Quantity:       promotion.FreeProductQuantity(li.Quantity),
					Type:           promotion.Name,
				}
				favorableLineItemsMap[promotion.Name] = append(favorableLineItemsMap[promotion.Name], favorableLineItem)
			}
		}
	}
	return favorableLineItemsMap
}

// Message returns the message of this Receipt.
// It will panic if template parsing fails
func (r Receipt) Message() string {
	tmpl, err := template.New("receipt.tmpl").Funcs(libs.ReceiptFuncMap).ParseFiles(libs.TmplFilePathWithName("receipt"))
	if err != nil {
		panic("Receipt#Message: " + err.Error())
	}

	var outputs bytes.Buffer
	tmpl.Execute(&outputs, r)
	// go1.5 does not suppot {{-
	return regexp.MustCompile("\n+").ReplaceAllString(outputs.String(), "\n")
}
