package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

func main() {
	// Floating point var for keeping track of time
	var time float32

	// Decimal floating point var for keeping track of time
	var decTime decimal.Decimal64

	// Given time over 100 hours
	for i := 0; i < 100*60*60*10; i++ {

		// Adding 0.1 repeatedly compounds rounding error
		time += 0.1                                              //s
		decTime = decTime.Add(decimal.MustParseDecimal64("0.1")) //s
	}

	// diff := decTime.Sub(decimal.MustParseDecimal64(fmt.Sprintf("%f", time))).Quo(decimal.MustParseDecimal64("10").Mul(decimal.MustParseDecimal64("60").Mul(decimal.MustParseDecimal64("60"))))
	diff := decTime.Sub(decimal.MustParseDecimal64(fmt.Sprintf("%f", time)))

	// time = time / (60 * 10 * 60)
	// time /= 10
	// foo := decimal.MustParseDecimal64("60").Mul(decimal.MustParseDecimal64("60")).Mul(decimal.MustParseDecimal64("10").Mul(decimal.MustParseDecimal64("100")))
	// decTime = decTime.Quo(decimal.MustParseDecimal64("10"))
	// 3600000
	// 36000
	// 360000
	// 3600000
	//
	fmt.Printf("Seconds taken in floating point: %f\nReal time taken in             : %f \nTime difference                : %f", time, decTime, diff)

}
