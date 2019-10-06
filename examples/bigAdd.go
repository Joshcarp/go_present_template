package Add

import "math/big"

func Add(x, y big.Float) big.Float {
	return *x.Add(&x, &y)
}
