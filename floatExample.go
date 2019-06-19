package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

// Simple example of how 1/10th cannot be truly be represented in a floating point number
func main() {
	var a = 0.1
	var b = 0.3
	fmt.Printf("%v + %v + %v == %v : %v\n", a, a, a, b, (a+a+a == b))

	// Printing out with 20 sig figs we can see where the problem is.
	fmt.Printf("%.20f + %.20f + %.20f == %.20f : %v\n", a, a, a, b, (a+a+a == b))

	// Using a decimal datatype can mitigate these errors as decimal fractions can be represented exactly
	c := decimal.MustParseDecimal64("0.1")
	d := decimal.MustParseDecimal64("0.3")

	fmt.Printf("%.20f + %.20f + %.20f == %.20f : %v\n", c, c, c, d, (c.Add(c).Add(c) == d))

}
