package decimal

import "fmt"

// Abs computes ||d||.
func (d *DecParts) Abs() *DecParts {
	d.sign = 0
	return d
}

// IsNaN computes d + e with default rounding
func (d *DecParts) IsNaN() bool {
	return d.fl == flQNaN || d.fl == flSNaN
}

// IsNaN computes d + e with default rounding
func (d *DecParts) String() string {
	return fmt.Sprintf("neg(%d) %d * 10 ^ %d", d.sign, d.significand, d.exp)
}

// Add computes d + e with default rounding
func (d *DecParts) Add(e *DecParts) *DecParts {
	return DefaultContext.DecAdd(d, e)
}

// FMA computes d*e + f with default rounding.
func (d *DecParts) FMA(e, f *DecParts) *DecParts {
	return DefaultContext.DecFMA(d, e, f)
}

// Mul computes d * e with default rounding.
func (d *DecParts) Mul(e *DecParts) *DecParts {
	return DefaultContext.DecMul(d, e)
}

// Sub returns d - e.
func (d *DecParts) Sub(e *DecParts) *DecParts {
	return d.Add(e.Neg())
}

// Quo computes d / e with default rounding.
func (d *DecParts) Quo(e *DecParts) *DecParts {
	return DefaultContext.DecQuo(d, e)
}

// Cmp returns:
//
//   -2 if d or e is NaN
//   -1 if d <  e
//    0 if d == e (incl. -0 == 0, -Inf == -Inf, and +Inf == +Inf)
//   +1 if d >  e
//
func (dp *DecParts) Cmp(ep *DecParts) int {
	if dec := propagateNan(dp, ep); dec != nil {
		return -2
	}
	if dp.isZero() && ep.isZero() {
		return 0
	}
	if dp.significand == ep.significand {
		return 0
	}
	dp = dp.Sub(ep)
	return 1 - 2*int(dp.significand.lo>>63)
}

// Neg computes -d.
func (d *DecParts) Neg() *DecParts {
	d.sign = ^d.sign
	return d
}

// DecQuo computes d / e.
func (ctx Context64) DecQuo(dp, ep *DecParts) *DecParts {
	var ans DecParts
	ans.sign = dp.sign ^ ep.sign
	if dp.isZero() {
		if ep.isZero() {
			return &DecNaN
		}
		return &DecZeroes[ans.sign]
	}
	if dp.isinf() {
		if ep.isinf() {
			return &DecNaN
		}
		return &DecInf
	}
	if ep.isinf() {
		return &DecZeroes[ans.sign]
	}
	if ep.isZero() {
		return &DecInf
	}
	dp.matchSignificandDigits(ep)
	ans.exp = dp.exp - ep.exp
	for {
		for dp.significand.gt(ep.significand) {
			dp.significand = dp.significand.sub(ep.significand)
			ans.significand = ans.significand.add(uint128T{1, 0})
		}
		if dp.significand == (uint128T{}) || ans.significand.numDecimalDigits() == 16 {
			break
		}
		ans.significand = ans.significand.mulBy10()
		dp.significand = dp.significand.mulBy10()
		ans.exp--
	}
	var rndStatus discardedDigit
	dp.significand = dp.significand.mul64(2)
	if dp.significand == (uint128T{}) {
		rndStatus = eq0
	} else if dp.significand.gt(ep.significand) {
		rndStatus = gt5
	} else if dp.significand.lt(ep.significand) {
		rndStatus = lt5
	} else {
		rndStatus = eq5
	}
	ans.significand.lo = ctx.roundingMode.round(ans.significand.lo, rndStatus)
	if ans.exp < -expOffset {
		rndStatus = ans.rescale(-expOffset)
		ans.significand.lo = ctx.roundingMode.round(ans.significand.lo, rndStatus)
	}
	if ans.exp >= -expOffset && ans.significand.lo != 0 {
		ans.exp, ans.significand.lo = renormalize(ans.exp, ans.significand.lo)
	}
	if ans.significand.lo > maxSig || ans.exp > expMax {
		return &DecInf
	}
	return &ans
}

// Sqrt computes √d.
// func (d DecParts) Sqrt() DecParts {
// 	flavor := d.fl
// 	sign := d.sign
// 	significand := d.significand.lo
// 	switch flavor {
// 	case flInf:
// 		if sign == 1 {
// 			return DecNaN
// 		}
// 		return d
// 	case flQNaN:
// 		return d
// 	case flSNaN:
// 		return DecNaN
// 	case flNormal:
// 	}
// 	if significand == 0 {
// 		return d
// 	}
// 	if sign == 1 {
// 		return DecNaN
// 	}
// 	if exp&1 == 1 {
// 		exp--
// 		significand *= 10
// 	}
// 	sqrt := umul64(10*DecPartsBase, significand).sqrt()
// 	exp, significand = renormalize(exp/2-8, sqrt)
// 	return DecParts{flNormal, sign, exp, significand}
// }

// DecAdd computes d + e
func (ctx Context64) DecAdd(dp, ep *DecParts) *DecParts {
	if dp.fl == flInf || ep.fl == flInf {
		if dp.fl != flInf {
			return ep
		}
		if ep.fl != flInf || ep.sign == dp.sign {
			return dp
		}
		return &DecNaN
	}
	if dp.significand.lo == 0 {
		return ep
	} else if ep.significand.lo == 0 {
		return dp
	}
	ep.removeZeros()
	dp.removeZeros()
	sep := dp.separation(ep)

	if sep < 0 {
		dp, ep = ep, dp
		sep = -sep
	}
	if sep > 17 {
		return dp
	}
	var rndStatus discardedDigit
	dp.matchScales128(ep)
	ans := dp.add128(ep)
	rndStatus = ans.roundToLo()
	if ans.exp < -expOffset {
		rndStatus = ans.rescale(-expOffset)
	}
	ans.significand.lo = ctx.roundingMode.round(ans.significand.lo, rndStatus)
	if ans.exp >= -expOffset && ans.significand.lo != 0 {
		ans.exp, ans.significand.lo = renormalize(ans.exp, ans.significand.lo)
	}
	if ans.exp > expMax || ans.significand.lo > maxSig {
		return &ans
	}
	return &ans
}

// DecFMA computes d*e + f
func (ctx Context64) DecFMA(dp, ep, fp *DecParts) *DecParts {
	var ans DecParts
	ans.sign = dp.sign ^ ep.sign
	if dp.fl == flInf || ep.fl == flInf {
		if fp.fl == flInf && ans.sign != fp.sign {
			return &DecNaN
		}
		if ep.isZero() || dp.isZero() {
			return &DecNaN
		}
		return &DecInf
	}
	if ep.significand.lo == 0 || dp.significand.lo == 0 {
		return fp
	}
	if fp.fl == flInf {
		return &DecInf
	}

	var rndStatus discardedDigit
	ep.removeZeros()
	dp.removeZeros()
	ans.exp = dp.exp + ep.exp
	ans.significand = umul64(dp.significand.lo, ep.significand.lo)
	sep := ans.separation(fp)
	if fp.significand.lo != 0 {
		if sep < -17 {
			return fp
		} else if sep <= 17 {
			ans = ans.add128(fp)
		}
	}
	rndStatus = ans.roundToLo()
	if ans.exp < -expOffset {
		rndStatus = ans.rescale(-expOffset)
	}
	ans.significand.lo = ctx.roundingMode.round(ans.significand.lo, rndStatus)
	if ans.exp >= -expOffset && ans.significand.lo != 0 {
		ans.exp, ans.significand.lo = renormalize(ans.exp, ans.significand.lo)
	}
	if ans.exp > expMax || ans.significand.lo > maxSig {
		return &DecInf
	}
	return &ans
}

// DecMul computes d * e
func (ctx Context64) DecMul(dp, ep *DecParts) *DecParts {

	var ans DecParts
	ans.sign = dp.sign ^ ep.sign
	if dp.fl == flInf || ep.fl == flInf {
		if ep.isZero() || dp.isZero() {
			return &DecNaN
		}
		return &DecInf
	}
	if ep.significand.lo == 0 || dp.significand.lo == 0 {
		return &DecZeroes[ans.sign]
	}
	var roundStatus discardedDigit
	significand := umul64(dp.significand.lo, ep.significand.lo)
	ans.exp = dp.exp + ep.exp + 15
	significand = significand.div64(decimal64Base)
	ans.significand.lo = significand.lo
	if ans.exp >= -expOffset {
		ans.exp, ans.significand.lo = renormalize(ans.exp, ans.significand.lo)
	} else if ans.exp < 1-expMax {
		roundStatus = ans.rescale(-expOffset)
	}
	ans.significand.lo = ctx.roundingMode.round(ans.significand.lo, roundStatus)
	if ans.significand.lo > maxSig || ans.exp > expMax {
		return &DecInf
	}
	return &ans
}
