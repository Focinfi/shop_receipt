package models

import (
	"fmt"
	"strings"

	"github.com/Focinfi/shop_receipt/libs"
)

type PromotionCounter interface {
	FreeProductQuantity(allProductQuantity int) int
	SubTotal(price float64, quantity int) float64
}

type PromotionType struct {
	Name string
	PromotionCounter
}

type ThreeForTwoPromotionCounter struct{}

func (counter ThreeForTwoPromotionCounter) FreeProductQuantity(allProductQuantity int) int {
	return allProductQuantity / 3
}

func (counter ThreeForTwoPromotionCounter) SubTotal(price float64, quantity int) float64 {
	left := quantity % 3
	return libs.Round(price*float64(quantity/3*2+left), 2)
}

type DiscountPromotionCounter struct{ Discount float64 }

func (counter DiscountPromotionCounter) FreeProductQuantity(allProductQuantity int) int {
	return 0
}

func (counter DiscountPromotionCounter) SubTotal(price float64, quantity int) float64 {
	return libs.Round(price*counter.Discount*float64(quantity), 2)
}

var ThreeForTwoPromotionType = PromotionType{Name: "3_for_2", PromotionCounter: ThreeForTwoPromotionCounter{}}

func MakeDiscountPromotionType(discount float64) PromotionType {
	if discount < 0 || discount >= 1 {
		panic("makeDiscountPromotionType: discount should be in (0~1)")
	}
	return PromotionType{Name: fmt.Sprintf("Discount_%-2f", discount), PromotionCounter: DiscountPromotionCounter{discount}}
}

type Promotion struct {
	PromotionType
	ProductBarCode string
	Weight         int
	Printable      bool
}

var promotions = map[string]Promotion{}

func FindAllPromotions(barCode string) []Promotion {
	all := []Promotion{}
	maxWeight := 0
	for key := range promotions {
		if strings.HasPrefix(key, barCode+"/") {
			if promotions[key].Weight > maxWeight {
				all = []Promotion{promotions[key]}
				maxWeight = promotions[key].Weight
			} else if promotions[key].Weight == maxWeight {
				all = append(all, promotions[key])
			}
		}
	}
	return all
}
