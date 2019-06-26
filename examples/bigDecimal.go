package main

import "math/big"

type Decimal struct {
	value *big.Int
	exp   int32
}
