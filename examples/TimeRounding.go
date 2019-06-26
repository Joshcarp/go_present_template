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
		time += 0.1
		decTime = decTime.Add(decimal.MustParseDecimal64("0.1"))
	}
	diff := decTime.Sub(decimal.MustParseDecimal64(fmt.Sprintf("%f", time)))
	fmt.Printf(`Seconds taken in floating point: %f
Real time taken in Decimal     : %f
Time difference                : %f`, time, decTime, diff)

}
