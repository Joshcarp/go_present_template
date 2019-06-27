package main

import (
	"fmt"
)

func main() {
	t := 60 * 60 * 100
	// Floating point var for keeping track of time
	var time float32
	// Decimal floating point var for keeping track of time

	for seconds := 0; seconds < t; seconds++ {
		// Adding 0.1 repeatedly compounds rounding error
		for tenths := 0; tenths < 10; tenths++ {
			time += 0.1
		}
	}
	diff := float32(t) - time
	fmt.Printf(`Seconds taken in floating point: %f s
Real time taken                : %d s
Time difference                :  %f s`, time, t, diff)

}
