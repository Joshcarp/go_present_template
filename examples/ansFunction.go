package main

// START OMIT
func (d Decimal64) Add(e Decimal64) Decimal64 {
	dp := d.getParts()
	ep := e.getParts()
	var ans DecParts
	// Arithmetic here ...
	return newFromParts(ans)
}

// END OMIT
