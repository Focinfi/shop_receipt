package libs

import (
	"fmt"
	"github.com/Focinfi/shop_receipt/config"
)

func MoneyOf(n float64) string {
	return fmt.Sprintf("%.2f(%s)", n, string(config.CurrentCurrency))
}

func ReceiptUnitKeyOf(unit string) string {
	return fmt.Sprintf("receipt_unit_%s", unit)
}

var ReceiptFuncMap = map[string]interface{}{
	"money_of":    MoneyOf,
	"unit_key_of": ReceiptUnitKeyOf,
	"t":           Translate,
}
