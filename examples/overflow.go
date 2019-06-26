package main

import "fmt"

func main() {
	var sig1 uint32 = 999999999
	var exp1 uint32 = 30
	var sig2 uint32 = 10
	var exp2 uint32 = 28

	exponentDifference := exp1 - exp2 // == 2

	ansSig := sig1*powersOf10[exponentDifference] + sig2
	ansExp := min(exp1, exp2)

	fmt.Printf("Answer = %d * 10 ^ %d", ansSig, ansExp)

}
func min(a ...uint32) uint32 {
	min := a[0]
	for _, val := range a {
		if val < min {
			min = val
		}
	}
	return min
}

var powersOf10 = []uint32{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
}
