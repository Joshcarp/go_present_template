package decimal

// DecParts stores the constituting DecParts of a decimal64.
// type DecParts struct {
// 	fl          flavor
// 	sign        int
// 	exp         int
// 	significand uint128T
// 	dec         *Decimal64
// }

var DecZero = DecParts{flNormal, 0, 0, uint128T{0, 0}, nil}
var NegDecZero = DecParts{flNormal, 1, 0, uint128T{0, 0}, nil}
var DecZeroes = []DecParts{DecZero, NegDecZero}
var DecInf = DecParts{flInf, 1, 0, uint128T{0, 0}, nil}
var DecNaN = DecParts{flQNaN, 1, 0, uint128T{0, 0}, nil}

// add128 adds two DecParts with full precision in 128 bits of significand
func (dp *DecParts) add128(ep *DecParts) DecParts {
	dp.matchScales128(ep)
	var ans DecParts
	ans.exp = dp.exp
	if dp.sign == ep.sign {
		ans.sign = dp.sign
		ans.significand = dp.significand.add(ep.significand)
	} else {
		if ep.significand.gt(dp.significand) {
			ans.sign = ep.sign
			ans.significand = ep.significand.sub(dp.significand)
		} else if ep.significand.lt(dp.significand) {
			ans.sign = dp.sign
			ans.significand = dp.significand.sub(ep.significand)
		} else {
			ans.significand = uint128T{0, 0}
		}
	}
	return ans
}

func (dp *DecParts) matchScales128(ep *DecParts) {
	expDiff := ep.exp - dp.exp
	if (ep.significand != uint128T{0, 0}) {
		if expDiff < 0 {
			dp.significand = dp.significand.mul(powerOfTen128(expDiff))
			dp.exp += expDiff
		} else if expDiff > 0 {
			ep.significand = ep.significand.mul(powerOfTen128(expDiff))
			ep.exp -= expDiff
		}
	}
}

func (dp *DecParts) matchSignificandDigits(ep *DecParts) {
	expDiff := ep.significand.numDecimalDigits() - dp.significand.numDecimalDigits()
	if expDiff >= 0 {
		dp.significand = dp.significand.mul(powerOfTen128(expDiff + 1))
		dp.exp -= expDiff + 1
		return
	}
	ep.significand = ep.significand.mul(powerOfTen128(-expDiff - 1))
	ep.exp -= -expDiff - 1
}

func (dp *DecParts) roundToLo() discardedDigit {
	var rndStatus discardedDigit
	if dp.significand.numDecimalDigits() > 16 {
		var remainder uint64
		expDiff := dp.significand.numDecimalDigits() - 16
		dp.exp += expDiff
		dp.significand, remainder = dp.significand.divrem64(powersOf10[expDiff])
		rndStatus = roundStatus(remainder, 0, expDiff)
	}
	return rndStatus
}

func (dp *DecParts) isZero() bool {
	return (dp.significand == uint128T{}) && dp.significand.hi == 0 && dp.fl == flNormal
}

func (dp *DecParts) isInf() bool {
	return dp.fl == flInf
}

func (dp *DecParts) isNaN() bool {
	return dp.fl == flQNaN || dp.fl == flSNaN
}

func (dp *DecParts) isQNaN() bool {
	return dp.fl == flQNaN
}

func (dp *DecParts) isSNaN() bool {
	return dp.fl == flSNaN
}

func (dp *DecParts) isSubnormal() bool {
	return (dp.significand != uint128T{}) && dp.significand.lo < decimal64Base && dp.fl == flNormal
}

// separation gets the separation in decimal places of the MSD's of two decimal 64s
func (dp *DecParts) separation(ep *DecParts) int {
	return dp.significand.numDecimalDigits() + dp.exp - ep.significand.numDecimalDigits() - ep.exp
}

// removeZeros removes zeros and increments the exponent to match.
func (dp *DecParts) removeZeros() {
	zeros := countTrailingZeros(dp.significand.lo)
	dp.significand.lo /= powersOf10[zeros]
	dp.exp += zeros
}

// isinf returns true if the decimal is an infinty
func (dp *DecParts) isinf() bool {
	return dp.fl == flInf
}

func (dp *DecParts) rescale(targetExp int) (rndStatus discardedDigit) {
	expDiff := targetExp - dp.exp
	mag := dp.significand.numDecimalDigits()
	rndStatus = roundStatus(dp.significand.lo, dp.exp, targetExp)
	if expDiff > mag {
		dp.significand.lo, dp.exp = 0, targetExp
		return
	}
	divisor := powersOf10[expDiff]
	dp.significand.lo = dp.significand.lo / divisor
	dp.exp = targetExp
	return
}
