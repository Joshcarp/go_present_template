package main

import (
	"fmt"
	"math"

	"github.com/anz-bank/decimal"
)

func main() {
	var P float64 = 100.1
	Pd := decimal.MustParseDecimal64("100.1")
	var r float64 = 0.045
	rd := decimal.MustParseDecimal64("0.045")
	var n int64 = 1

	for i := float64(1); i < 100; i++ {

		P = compoundSimple(P, r, n)
		Pd = compoundDecimal(Pd, rd, n)
		fmt.Println(i, P, Pd)

	}

	// var Pd = decimal.MustParseDecimal64("100.1")
	// var Pd = decimal.MustParseDecimal64("0.045")
	// var Pd = decimal.MustParseDecimal64("12")

}

// func compound(P float64, r float64, n int, t float64) (A float64) {
// 	return P * math.Pow((1+(r/n)), float64(n)*t)
//
// }
func compoundSimple(P float64, r float64, n int64) (A float64) {
	return P * math.Pow((1+(r/float64(n))), float64(n))

}

func compoundDecimal(P decimal.Decimal64, r decimal.Decimal64, n int64) (A decimal.Decimal64) {

	return P.Mul(decPow((decimal.NewDecimal64FromInt64(1)).Add((r.Quo(decimal.NewDecimal64FromInt64(n)))), n))

}

func decPow(d decimal.Decimal64, p int64) decimal.Decimal64 {
	return d
	// original := d
	// for i := p; i >= 0; i-- {
	// 	d = d.Mul(original)
	// }
	// return d
}

// func compoundDecimal(P decimal.Decimal64, r decimal.Decimal64, n decimal.Decimal64, t decimal.Decimal64) (A decimal.Decimal64) {
// 	return P * math.Pow((1+(r/n)), n*t)

// }

// P -> Principal, dollars, float64
// r -> rate, percentage, float64
// t -> time, years, int
// n -> number of installments per unit t
// A -> amount after time
