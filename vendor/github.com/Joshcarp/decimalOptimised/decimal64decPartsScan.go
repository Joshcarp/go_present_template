package decimal

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

// ParseDecParts parses a string representation of a number as a Decimal.
func ParseDecParts(s string) (*DecParts, error) {
	state := stringScanner{reader: strings.NewReader(s)}
	var d DecParts
	if err := d.Scan(&state, 'e'); err != nil {
		return &d, err
	}

	// entire string must have been consumed
	r, _, err := state.ReadRune()
	if err == nil {
		return &d, fmt.Errorf("expected end of string, found %c", r)
	}
	logicCheck(err == io.EOF, "%v == io.EOF", err)
	return &d, nil
}

// MustParseDecParts parses a string as a Decimal and returns the value or
// panics if the string doesn't represent a valid Decimal.
func MustParseDecParts(s string) DecParts {
	d, err := ParseDecParts(s)
	if err != nil {
		panic(err)
	}
	return *d
}

// Scan implements fmt.Scanner.
func (d *DecParts) Scan(state fmt.ScanState, verb rune) error {
	// *d = SNaN

	sign, err := scanSign(state)
	if err != nil {
		return err
	}
	// Word-number ([Ii]nf|∞|nan|NaN)
	word, err := tokenString(state, isLetterOrInf)
	if err != nil {
		return err
	}
	switch strings.ToLower(word) {
	case "":
	case "inf", "infinity", "∞":
		if sign == 0 {
			*d = DecInf
		} else {
			*d = DecInf
		}
		return nil
	case "nan", "qnan":
		payload, _ := tokenString(state, unicode.IsDigit)
		payloadInt, _ := parseUint(payload)
		*d = newDecPayloadNan(sign, flQNaN, uint64(payloadInt))
		*d = DecNaN
		return nil
	case "snan":
		payload, _ := tokenString(state, unicode.IsDigit)
		payloadInt, _ := parseUint(payload)
		*d = newDecPayloadNan(sign, flSNaN, uint64(payloadInt))
		*d = DecNaN
		return nil
	default:
		return notDecimal64()
	}

	whole, err := tokenString(state, unicode.IsDigit)
	if err != nil {
		return err
	}

	dot, err := tokenString(state, func(r rune) bool { return r == '.' })
	if err != nil {
		return err
	}
	if len(dot) > 1 {
		return fmt.Errorf("Too many dots")
	}

	frac, err := tokenString(state, unicode.IsDigit)
	if err != nil {
		return err
	}

	e, err := tokenString(state, func(r rune) bool { return r == 'e' || r == 'E' })
	if err != nil {
		return err
	}
	if len(e) > 1 {
		return fmt.Errorf("Too many 'e's")
	}

	var expSign int
	var exp string
	if len(e) == 1 {
		expSign, err = scanSign(state)
		if err != nil {
			return err
		}
		exp, err = tokenString(state, unicode.IsDigit)
		if err != nil {
			return err
		}
		if exp == "" {
			return fmt.Errorf("Exponent value missing")
		}
	}

	mantissa := whole + frac
	if mantissa == "" {
		return fmt.Errorf("Mantissa missing")
	}
	mantissa = strings.TrimLeft(mantissa, "0")
	if mantissa == "" {
		mantissa = "0"
	}

	significand, sExp := parseUint(mantissa)
	if significand == 0 {
		*d = DecZero
		return nil
	}

	exponent, _ := parseUint(exp)
	exponent *= int64(1 - 2*expSign)
	if exponent > 1000 {
		*d = DecZero
		return nil
	} else if exponent < -1000 {
		*d = DecZeroes[sign]
		return nil
	}
	exponent += int64(sExp - len(frac))

	partExp, partSignificand := renormalize(int(exponent), uint64(significand))
	// *d = newFromParts(sign, partExp, partSignificand)
	*d = DecParts{flNormal, sign, partExp, uint128T{0, partSignificand}, nil}
	return nil
}

func newDecPayloadNan(sign int, fl flavor, weight uint64) DecParts {
	switch fl {
	case flQNaN:
		return DecParts{flQNaN, sign, 0, uint128T{0, weight}, nil}
	case flSNaN:
		return DecParts{flSNaN, sign, 0, uint128T{0, weight}, nil}
	default:
		return DecParts{flQNaN, sign, 0, uint128T{0, weight}, nil}
	}
}
