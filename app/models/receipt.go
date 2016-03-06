package models

import (
	"fmt"
	"github.com/Focinfi/shop_receipt/config"
	"regexp"
	"strconv"
	"strings"
)

type Receipt struct {
	LineItems map[string]*LineItem
}

func NewReceipt(barCodes []string) *Receipt {
	lineItems := map[string]*LineItem{}
	// init lineItems
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

func (r Receipt) Total() float64 {
	var total float64
	for _, li := range r.LineItems {
		total += li.Subtotal()
		fmt.Println(li)
	}
	return total
}

func (r Receipt) CostSaving() float64 {
	var costSaving float64
	for _, li := range r.LineItems {
		costSaving += li.CostSaving()
	}
	return costSaving
}

func (r Receipt) SpecialOffers() map[string][]FavorableLineItem {
	favorableLineItemsMap := map[string][]FavorableLineItem{}
	for _, li := range r.LineItems {
		for _, promotion := range FindAllPromotions(li.product.BarCode) {
			if promotion.Printable {
				favorableLineItem := FavorableLineItem{
					ProductBarCode: li.product.BarCode,
					Quantity:       promotion.FreeProductQuantity(li.Quantity),
					Type:           promotion.Name,
				}
				favorableLineItemsMap[promotion.Name] = append(favorableLineItemsMap[promotion.Name], favorableLineItem)
			}
		}
	}
	return favorableLineItemsMap
}

func (r Receipt) Message() string {
	header := fmt.Sprintf("******<%s> shopping list\n", config.AppName)
	footer := "************************\n"
	intervalLine := "------------------------\n"
	var lineItems string
	var favorableLineItems string
	var total string

	for _, li := range r.LineItems {
		lineItems += fmt.Sprintf("%v\n", li)
	}

	for name, lineItems := range r.SpecialOffers() {
		favorableLineItems += fmt.Sprintf("%s:\n", name)
		for _, li := range lineItems {
			favorableLineItems += fmt.Sprintf("%v\n", li)
		}
	}

	total = fmt.Sprintf("total: %.2f \n", r.Total())
	return fmt.Sprint(header, lineItems, intervalLine, favorableLineItems, intervalLine, total, footer)
}
