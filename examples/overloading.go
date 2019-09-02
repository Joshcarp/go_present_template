package main

import "github.com/anz-bank/decimal"

func main() {
	// START OMIT
	b := decimal.MustParseDecimal64("0.3")
	c := decimal.MustParseDecimal64("0.1")
	a := c.Mul(a)

	// or something like
	a := 0.3
	b := 0.1
	c := c * a

	// END OMIT

}
