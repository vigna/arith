package frac

import "strconv"

// Creation
func NewFraction(num, denom int64) Fraction {
	if gcd := gcd(num, denom); gcd != 1 {
		num /= gcd
		denom /= gcd
	}
	return Fraction{num, denom}
}

func (f Fraction) String() string {
	if f.denom == 1 {
		return strconv.Itoa(int(f.num))
	}
	return strconv.Itoa(int(f.num)) + "/" + strconv.Itoa(int(f.denom))
}