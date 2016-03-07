package libs

import (
	"fmt"
	"github.com/Focinfi/shop_receipt/config"
	"strings"
)

func Translate(key string) string {
	translationKey := fmt.Sprintf("%s_%s", key, config.CurrentLocal)
	translated, ok := config.Translations[translationKey]
	if !ok {
		translated = strings.Replace(key, "_", " ", -1)
	}
	return translated
}
