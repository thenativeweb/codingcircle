package egyptianfractions

func Reduce(fraction Fraction) Fraction {
	divisor := GCD(fraction.Numerator, fraction.Denominator)

	return Fraction{
		Numerator:   fraction.Numerator / divisor,
		Denominator: fraction.Denominator / divisor,
	}
}
