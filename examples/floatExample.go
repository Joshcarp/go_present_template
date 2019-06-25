package main

import (
	"fmt"
)

func main() {
	a := 0.1
	b := 0.3

	fmt.Printf("%f, %f, %f, %f, 3*a == b; %v\n", a, a, a, b, 3*a == b)

	fmt.Printf("%.20f, %.20f, %.20f, %.20f, 3*a == b; %v\n", a, a, a, b, 3*a == b)

}
