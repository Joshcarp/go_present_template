package main

import (
	"fmt"
	"math"
)

// Simple example of how 1/10th cannot be truly be represented in a floating point number
func main() {
	var a = 0.01
	fmt.Printf("%.20f\n", a)
	// for i := 0; i < 100; i++ {
	i := 99
	b := float64(1) / math.Pow(float64(2), float64(i))
	fmt.Printf("%d, %.200f\n", i, b)
	// }

}
