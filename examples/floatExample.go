package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

func main() {
	var a decimal.Decimal64
	b := decimal.MustParseDecimal64("0.1")
	c := decimal.MustParseDecimal64("0.3")
	d := decimal.NewDecimal64FromInt64(123456)

	fmt.Println(a, b, c, d)

}
