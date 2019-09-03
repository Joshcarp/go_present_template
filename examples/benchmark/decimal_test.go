package main

import (
	"testing"

	"github.com/anz-bank/decimal"
)

func BenchmarkANZ(b *testing.B) {
	// run the Fib function b.N times
	var a, c decimal.Decimal64

	for n := 0; n < 1000000; n++ {
		a.Add(c)
	}
}
