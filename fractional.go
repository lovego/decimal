package decimal

import "math/big"

// FractionalBits returns the number of significant fractional digits.
func (d Decimal) FractionalBits() uint {
	if d.Equal(Zero) || d.exp >= 0 {
		return 0
	}
	exp := uint(-d.exp)
	zeroBits := TrailingZeroBits(d.value) // only go1.14+ big.Int have TrailingZeroBits method
	if exp > zeroBits {
		return exp - zeroBits
	}
	return 0
}

var intTen = big.NewInt(10)
var intZero = big.NewInt(0)

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
