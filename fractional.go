package decimal

// FractionalBits returns the number of significant fractional digits.
func (d Decimal) FractionalBits() uint {
	if d.Equal(Zero) || d.exp >= 0 {
		return 0
	}
	exp := uint(-d.exp)
	zeroBits := d.value.TrailingZeroBits()
	if exp > zeroBits {
		return exp - zeroBits
	}
	return 0
}
