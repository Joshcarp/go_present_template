package main

import (
	"fmt"
)

func main() {
	t := 60 * 60 * 100
	var time float32

	for seconds := 0; seconds < t; seconds++ {

		for tenths := 0; tenths < 10; tenths++ {
			time += 0.1
		}
	}
	diff := float32(t) - time
	fmt.Printf(`Seconds taken in floating point: %f s
Real time taken                : %d s
Time difference                :  %f s`, time, t, diff)

}
