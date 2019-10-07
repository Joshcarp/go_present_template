package main

import (
	"fmt"

	"github.com/anz-bank/decimal"
)

func main() {
	var t int64 = 60 * 60 * 100
	var seconds int64
	time := decimal.Decimal64{}

	for seconds = 0; seconds < t; seconds++ {

		for tenths := 0; tenths < 10; tenths++ {
			time = time.Add(decimal.MustParseDecimal64("0.1"))
		}
	}
	diff := decimal.NewDecimal64FromInt64(t).Sub(time)
	fmt.Printf(`Seconds taken in floating point: %f s
Real time taken                : %d s
Time difference                :  %f s`, time, t, diff)

}
