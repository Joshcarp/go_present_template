package main

import "math/big"

func Add(x, y, z big.Float) *big.Float {
	return x.Add(&y, &z)
}
