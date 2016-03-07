package models

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// For more readable
var Describ = Convey
var Context = Convey
var It = Convey
var Expect = So

var (
	// Product
	cola         = Product{BarCode: "ITEM000005", Name: "可口可乐", Uint: "瓶", Price: float64(3)}
	badminton    = Product{BarCode: "ITEM000001", Name: "羽毛球", Uint: "个", Price: float64(1)}
	apple        = Product{BarCode: "ITEM000003", Name: "苹果", Uint: "斤", Price: float64(5.5)}
	testProducts = map[string]Product{badminton.BarCode: badminton, apple.BarCode: apple, cola.BarCode: cola}

	// Promotion
	badminton3For2Promotion              = Promotion{ThreeForTwoPromotionType, badminton.BarCode, 99, true}
	cola3For2Promotion                   = Promotion{ThreeForTwoPromotionType, cola.BarCode, 99, true}
	badmintonNinetyFiveDiscountPromotion = Promotion{MakeDiscountPromotionType(0.95), badminton.BarCode, 10, false}
	appleNineFiveDiscountPromotion       = Promotion{MakeDiscountPromotionType(0.95), apple.BarCode, 10, false}

	testPromotions = map[string]Promotion{
		fmt.Sprintf("%s/%s", badminton.BarCode, badminton3For2Promotion.Name):              badminton3For2Promotion,
		fmt.Sprintf("%s/%s", badminton.BarCode, badmintonNinetyFiveDiscountPromotion.Name): badmintonNinetyFiveDiscountPromotion,
		fmt.Sprintf("%s/%s", apple.BarCode, appleNineFiveDiscountPromotion.Name):           appleNineFiveDiscountPromotion,
	}

	testBarCodes = []string{
		badminton.BarCode,
		badminton.BarCode,
		badminton.BarCode,
		apple.BarCode,
		fmt.Sprintf("%s-%d", cola.BarCode, 2),
	}

	demoBarCodes = []string{
		"ITEM000001",
		"ITEM000001",
		"ITEM000001",
		"ITEM000001",
		"ITEM000001",
		"ITEM000003-2",
		"ITEM000005",
		"ITEM000005",
		"ITEM000005",
	}
)

func TestLineItem(t *testing.T) {
	Describ("NewLineItem", t, func() {
		products = testProducts
		promotions = testPromotions
		lineItem := NewLineItem(badminton.BarCode, 3)
		It("creates a new LineItem", func() {
			Expect(lineItem.Subtotal(), ShouldEqual, float64(2))
			Expect(lineItem.CostSaving(), ShouldEqual, float64(1))
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
			Expect(receipt.Total(), ShouldEqual, float64(13.23))
			Expect(receipt.CostSaving(), ShouldEqual, float64(1.27))
			Expect(receipt.FavorableLineItemMap()["3_for_2"][0].Quantity, ShouldEqual, 1)
		})
	})
}

func TestMessage(t *testing.T) {
	Describ("Message", t, func() {
		Context("when there has promotion 3 for 2", func() {
			products = testProducts
			promotions = map[string]Promotion{
				fmt.Sprintf("%s/%s", badminton.BarCode, badminton3For2Promotion.Name): badminton3For2Promotion,
				fmt.Sprintf("%s/%s", cola.BarCode, cola3For2Promotion.Name):           cola3For2Promotion,
			}
			t.Logf("\n%s", NewReceipt(demoBarCodes).Message())
		})

		Context("when there has no promotions", func() {
			products = testProducts
			promotions = map[string]Promotion{}
			t.Logf("\n%s", NewReceipt(demoBarCodes).Message())
		})

		Context("when there has 95 persent discount no promotions", func() {
			products = testProducts
			promotions = map[string]Promotion{
				fmt.Sprintf("%s/%s", apple.BarCode, appleNineFiveDiscountPromotion.Name): appleNineFiveDiscountPromotion,
			}
			t.Logf("\n%s", NewReceipt(demoBarCodes).Message())
		})

		Context("when there has 3 for 2 and 95 persent discount promotion ", func() {
			products = testProducts
			promotions = map[string]Promotion{
				fmt.Sprintf("%s/%s", badminton.BarCode, badminton3For2Promotion.Name):    badminton3For2Promotion,
				fmt.Sprintf("%s/%s", cola.BarCode, cola3For2Promotion.Name):              cola3For2Promotion,
				fmt.Sprintf("%s/%s", apple.BarCode, appleNineFiveDiscountPromotion.Name): appleNineFiveDiscountPromotion,
			}
			t.Logf("\n%s", NewReceipt(append(demoBarCodes, "ITEM000001")).Message())
		})
	})
}
