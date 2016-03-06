package libs

import (
	"math"
	"os"
	"path"
)

func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

func TmplFilePathWithName(name string) string {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		panic("TmplWithName: has no GOPATH env")
	}
	return path.Join(goPath, "src", "github.com", "Focinfi", "shop_receipt", "app", "views", name+".tmpl")
}
