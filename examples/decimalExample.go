package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

// Simple example of how 1/10th cannot be truly be represented in a floating point number
func main() {
	a := decimal.MustParseDecimal64("0.1")
	b := decimal.MustParseDecimal64("0.3")
	c := decimal.NewDecimal64FromInt64(3)

	fmt.Printf("%.1f + %.1f + %.1f == %.1f; %v\n", a, a, a, b, c.Mul(a) == b)

}
