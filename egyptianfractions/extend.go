package egyptianfractions

func Extend(a, b Fraction) (Fraction, Fraction) {
	denominator := a.Denominator * b.Denominator

	extendedA := Fraction{
		Numerator:   a.Numerator * b.Denominator,
		Denominator: denominator,
	}
	extendedB := Fraction{
		Numerator:   b.Numerator * a.Denominator,
		Denominator: denominator,
	}

	return extendedA, extendedB
}
