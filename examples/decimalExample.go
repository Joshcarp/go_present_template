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

	fmt.Printf("%f, %f, %f, %f, 3*a == b; %v\n", a, a, a, b, a.Mul(c) == b)

}
