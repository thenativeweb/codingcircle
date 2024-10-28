package egyptianfractions

import "math"

func FindEgyptianFractions(fraction Fraction) []Fraction {
	fractions := []Fraction{}

	for fraction.Numerator > 0 {
		n := int(math.Ceil(
			float64(fraction.Denominator) /
				float64(fraction.Numerator)))

		partialFraction := Fraction{1, n}

		fractions = append(fractions, partialFraction)

		extendedFraction, extendedPartialFraction :=
			Extend(fraction, partialFraction)

		fraction = Reduce(Fraction{
			Numerator:   extendedFraction.Numerator - extendedPartialFraction.Numerator,
			Denominator: extendedFraction.Denominator,
		})
	}

	return fractions
}
