package frac

func (f Fraction) Mult(g Fraction) (r Fraction) {
	r.num = f.num * g.num
	r.denom = f.denom * g.denom
	r.reduce()
	return
}

func (f Fraction) Add(g Fraction) (r Fraction) {
	mcm := f.denom * g.denom / gcd(f.denom, g.denom)
	r.num = f.num * (mcm / f.denom) + g.num * (mcm / g.denom)
	r.denom = mcm
	r.reduce()
	return
}

func gcd(a, b int64) int64 {
	if a % b == 0 {
		return b
	}
	return gcd(b, a % b)
}

func (f *Fraction) reduce() {
	if gcd := gcd(f.num, f.denom); gcd != 1 {
		f.num /= gcd
		f.denom /= gcd
	}
}
