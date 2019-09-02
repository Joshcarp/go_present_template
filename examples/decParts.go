package main

// Decimal64 represents an IEEE 754 64-bit floating point decimal number.
// It uses the binary representation method.
type Decimal64 struct {
	bits uint64
}

// DecParts stores the constituting DecParts of a decimal64.
type DecParts struct {
	fl          flavor
	sign        int
	exp         int
	significand uint128T
	dec         *Decimal64
}
