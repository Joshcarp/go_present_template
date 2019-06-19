package main

//
// import (
// 	"fmt"
// 	"math"
//
// 	"github.com/shopspring/decimal"
// )
//
// func main() {
// 	var P float64 = 100.1
// 	Pd, _ := decimal.NewFromString("100.1")
// 	var r float64 = 0.045
// 	rd, _ := decimal.NewFromString("0.045")
// 	var n int64 = 1
//
// 	for i := float64(1); i < 500; i++ {
//
// 		P = compoundSimple(P, r, n)
// 		Pd = compoundDecimal(Pd, rd, n)
// 		fmt.Println(i, P, Pd)
//
// 	}
//     // 3.463083878119739
//     // 3.463083878119856063387653
// 	// var Pd = decimal.MustParseDecimal("100.1")
// 	// var Pd = decimal.MustParseDecimal("0.045")
// 	// var Pd = decimal.MustParseDecimal("12")
//
// }
//
// // func compound(P float64, r float64, n int, t float64) (A float64) {
// // 	return P * math.Pow((1+(r/n)), float64(n)*t)
// //
// // }
// func compoundSimple(P float64, r float64, n int64) (A float64) {
// 	return P * math.Pow((1+(r/float64(n))), float64(1))
//
// }
//
// func compoundDecimal(P decimal.Decimal, r decimal.Decimal, n int64) (A decimal.Decimal) {
// 	one, _ := decimal.NewFromString("1")
// 	nd, _ := decimal.NewFromString(fmt.Sprintf("%d", 1))
//
// 	return P.Mul(one.Add(r.Div(nd)).Pow(nd))
//
// }
//
// // func decPow(d decimal.Decimal, p int64) decimal.Decimal {
// // 	original := d
// // 	for i := p; i >= 0; i-- {
// // 		d = d.Mul(original)
// // 	}
// // 	return d
// // }
//
// // func compoundDecimal(P decimal.Decimal, r decimal.Decimal, n decimal.Decimal, t decimal.Decimal) (A decimal.Decimal) {
// // 	return P * math.Pow((1+(r/n)), n*t)
//
// // }
//
// // P -> Principal, dollars, float64
// // r -> rate, percentage, float64
// // t -> time, years, int
// // n -> number of installments per unit t
// // A -> amount after time
