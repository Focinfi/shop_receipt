package libs

import (
	"fmt"
	"github.com/Focinfi/shop_receipt/config"
)

func MoneyOf(n float64) string {
	return fmt.Sprintf("%.2f(%s)", n, string(config.CurrentCurrency))
}

var ReceiptFuncMap = map[string]interface{}{
	"money_of": MoneyOf,
	"t":        Translation,
}
