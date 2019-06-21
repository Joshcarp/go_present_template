package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

// Simple example of how 1/10th cannot be truly be represented in a floating point number
func main() {
	a := 0.1
	b := 0.3
	fmt.Println("Float datatype:")
	fmt.Printf("%v + %v + %v == %v : %v\n", a, a, a, b, (a+a+a == b))
	// // Printing out with 20 sig figs we can see where the problem is.
	fmt.Printf("%.20f + %.20f + %.20f == %.20f : %v\n\n", a, a, a, b, (a+a+a == b))
	//
	// Using a decimal datatype can mitigate these errors as decimal fractions can be represented exactly
	fmt.Println("Decimal datatype:")
	c := decimal.MustParseDecimal64("0.1")
	d := decimal.MustParseDecimal64("0.3")
	//
	fmt.Printf("%.20f + %.20f + %.20f == %.20f : %v\n\n", c, c, c, d, (c.Add(c).Add(c) == d))

}
