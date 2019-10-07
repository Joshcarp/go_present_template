package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

func main() {
	a := decimal.MustParseDecimal64("0.1")
	b := decimal.MustParseDecimal64("0.3")

	three := decimal.MustParseDecimal64("3")

	fmt.Printf("3*a == b; %v\n", three.Mul(a) == b)

	fmt.Printf("%.20f %.20f\n", a, b)

}
