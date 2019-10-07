package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.49999999999999999
	fmt.Println(math.Round(x))
	fmt.Printf("%.20f\n", x)
}
