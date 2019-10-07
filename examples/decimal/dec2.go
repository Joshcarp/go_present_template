package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

func main() {

	a := decimal.NewDecimal64FromInt64(10000000000000000)
	b := decimal.NewDecimal64FromInt64(1)
	ten := decimal.NewDecimal64FromInt64(10)

	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a.Add(b))
	fmt.Printf("%.2f + %.2f = %.2f\n", a, ten, a.Add(ten))
}
