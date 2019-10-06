package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

func main() {
	a := decimal.MustParseDecimal64("0.1")
	b := decimal.MustParseDecimal64("0.3")
	c := decimal.NewDecimal64FromInt64(3)

	fmt.Printf("%.1f + %.1f + %.1f == %.1f; %v\n", a, a, a, b, c.Mul(a) == b)

}
