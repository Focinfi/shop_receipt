package models

import (
	"fmt"
	"strings"

	"github.com/Focinfi/shop_receipt/libs"
)

// PromotionCounter defines two behavior of a promotion counter
type PromotionCounter interface {
	// FreeProductQuantity should return how many products can be free
	FreeProductQuantity(allProductQuantity int) int
	// SubTotal should return the subtotal
	SubTotal(price float64, quantity int) float64
}

type PromotionType struct {
	Name string
	PromotionCounter
}

// ThreeForTwoPromotionCounter defines for 3 for 2 promotion
type ThreeForTwoPromotionCounter struct{}

func (counter ThreeForTwoPromotionCounter) FreeProductQuantity(allProductQuantity int) int {
	return allProductQuantity / 3
}

func (counter ThreeForTwoPromotionCounter) SubTotal(price float64, quantity int) float64 {
	left := quantity % 3
	return libs.Round(price*float64(quantity/3*2+left), 2)
}

// DiscountPromotionCounter defines for discount poromtion
type DiscountPromotionCounter struct{ Discount float64 }

func (counter DiscountPromotionCounter) FreeProductQuantity(allProductQuantity int) int {
	return 0
}

func (counter DiscountPromotionCounter) SubTotal(price float64, quantity int) float64 {
	return libs.Round(price*counter.Discount*float64(quantity), 2)
}

var ThreeForTwoPromotionType = PromotionType{Name: "3_for_2", PromotionCounter: ThreeForTwoPromotionCounter{}}

// MakeDiscountPromotionType allocates and returns a discount promotion type with the given discount float64
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
	// Printable signs if this promotion should be printed in favorable products part of receipt
	Printable bool
}

var promotions = map[string]Promotion{}

// FindAllPromotions finds and return all promotions as a []Promotion of a product with the given barCode
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
