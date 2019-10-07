package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

func main() {
	a := decimal.MustParseDecimal64("0.49999999999999999")
	fmt.Println(a)
}
