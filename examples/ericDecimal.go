package main

import "math/big"

// START OMIT
type Big struct {
	// Context is used to give info like rounding modes and all
	Context Context

	// unscaled is only used if the decimal is too large to fit in compact.
	unscaled big.Int
	compact  uint64
	exp      int

	// ... more utility fields ...
	// END OMIT

}

// // START OMIT2

// func (c Context) Add(z, x, y *Big) *Big {

// 	// END OMIT2
// }
