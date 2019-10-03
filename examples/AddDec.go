package Decimal

// START OMIT
func (d Decimal64) Add(e Decimal64) Decimal64 {
	dp := d.getParts()
	ep := e.getParts()
	if dec := propagateNan(&dp, &ep); dec != nil {
		return *dec
	}
	if dp.fl == flInf || ep.fl == flInf {
		// Check infinities
	}
	// ...
	// Match scales
	var rndStatus discardedDigit
	dp.matchScales128(&ep)
	ans := dp.add128(&ep)
	rndStatus = ans.roundToLo()
	if ans.exp < -expOffset {
		rndStatus = ans.rescale(-expOffset)
	}
	// Do the calculation ...
	ans.significand.lo = ctx.roundingMode.round(ans.significand.lo, rndStatus)
	if ans.exp >= -expOffset && ans.significand.lo != 0 {
		ans.exp, ans.significand.lo = renormalize(ans.exp, ans.significand.lo)
	}
	if ans.exp > expMax || ans.significand.lo > maxSig {
		return infinities[ans.sign]
	}
	return newFromParts(ans.sign, ans.exp, ans.significand.lo)
}

// END OMIT
