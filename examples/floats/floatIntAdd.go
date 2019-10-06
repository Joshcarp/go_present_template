package main

import "fmt"

func main() {
	var x float32 = 33554430
	var y float32 = 4
	fmt.Printf("%.2f + %.2f = %.2f", x, y, x+y)
}
