package main

import (
	"fmt"
)

func main() {
	a := 0.1
	b := 0.3

	fmt.Printf("%.1f + %.1f + %.1f == %.1f; %v\n", a, a, a, b, 3*a == b)

	fmt.Printf("%.20f %.20f\n", a, b)

}
