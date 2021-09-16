package decimal

import "math/big"

func (d Decimal) IntBits() uint {
	var value = d.IntPart()
	if value < 0 {
		value = -value
	}
	var bits uint = 1
	for {
		value = value / 10
		if value > 0 {
			bits++
		} else {
			break
		}
	}
	return bits
}

// FractionalBits returns the number of significant fractional digits.
func (d Decimal) FractionalBits() uint {
	if d.Equal(Zero) || d.exp >= 0 {
		return 0
	}
	exponent := uint(-d.exp)
	zeroBits := TrailingZeroBits(d.value) // only go1.14+ big.Int have TrailingZeroBits method
	if exponent > zeroBits {
		return exponent - zeroBits
	}
	return 0
}

var intTen = big.NewInt(10)
var intZero = big.NewInt(0)

// how many zero at end of value
func TrailingZeroBits(value *big.Int) uint {
	var quotient, remaider = big.NewInt(0).Set(value), big.NewInt(0)
	var bits uint
	for {
		quotient.DivMod(quotient, intTen, remaider)
		if quotient.Cmp(intZero) != 0 && remaider.Cmp(intZero) == 0 {
			bits++
		} else {
			break
		}
	}
	return bits
}
