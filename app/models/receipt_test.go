package models

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var Describ = Convey
var Context = Convey
var It = Convey
var Expect = So

var (
	cup          = Product{BarCode: "FTC9001001", Name: "Cup", Price: float64(14), Currency: RMB}
	coffee       = Product{BarCode: "FTC9001002", Name: "Coffee", Price: float64(4), Currency: RMB}
	sugar        = Product{BarCode: "FTC9001003", Name: "Sugar", Price: float64(0.5), Currency: RMB}
	testProducts = map[string]Product{cup.BarCode: cup, coffee.BarCode: coffee, sugar.BarCode: sugar}

	cup3For2Promotion              = Promotion{ThreeForTwoPromotionType, cup.BarCode, 99, true}
	cupNinetyFiveDiscountPromotion = Promotion{MakeDiscountPromotionType(0.95), cup.BarCode, 10, false}
	coffeeNinetyDiscountPromotion  = Promotion{MakeDiscountPromotionType(0.90), coffee.BarCode, 10, false}

	testPromotions = map[string]Promotion{
		fmt.Sprintf("%s/%s", cup.BarCode, cup3For2Promotion.Name):                cup3For2Promotion,
		fmt.Sprintf("%s/%s", cup.BarCode, cupNinetyFiveDiscountPromotion.Name):   cupNinetyFiveDiscountPromotion,
		fmt.Sprintf("%s/%s", coffee.BarCode, coffeeNinetyDiscountPromotion.Name): coffeeNinetyDiscountPromotion,
	}
	testBarCodes = []string{
		cup.BarCode,
		cup.BarCode,
		cup.BarCode,
		coffee.BarCode,
		fmt.Sprintf("%s-%d", sugar.BarCode, 2),
	}
)

func TestLineItem(t *testing.T) {
	Describ("NewLineItem", t, func() {
		products = testProducts
		promotions = testPromotions
		lineItem := NewLineItem(cup.BarCode, 3)
		It("creates a new LineItem", func() {
			Expect(lineItem.Subtotal(), ShouldEqual, float64(28))
			Expect(lineItem.CostSaving(), ShouldEqual, float64(14))
			Expect(lineItem.String(), ShouldEqual, "name: Cup, quantity: 3, unit price: 14.00(元), subtotal: 28.00(元), cost saving: 14.00(元)")
		})
	})
}

func TestReceipt(t *testing.T) {
	Describ("NewReceipt", t, func() {
		products = testProducts
		promotions = testPromotions
		receipt := NewReceipt(testBarCodes)
		It("creates a new Receipt", func() {
			Expect(len(receipt.LineItems), ShouldEqual, 3)
			Expect(receipt.Total(), ShouldEqual, float64(14*2+4*0.9+0.5*2))
			Expect(receipt.CostSaving(), ShouldEqual, float64(14*1+4*0.1))
			Expect(receipt.SpecialOffers()["3 for 2"][0].Quantity, ShouldEqual, 1)
		})
	})
}
