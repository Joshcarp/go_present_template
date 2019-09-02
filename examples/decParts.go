package main

// START OMIT
type Decimal64 struct {
	bits uint64
}

type DecParts struct {
	fl          flavor
	sign        int
	exp         int
	significand uint128T
	dec         *Decimal64
}

// END OMIT
