package clevermax

import "math"

func abs(i int) int {
	return int(math.Sqrt(float64(i * i)))
}

func Max(left, right int) int {
	return (left + right + abs(left-right)) / 2
}

func MaxBitShift(left, right int) int {
	difference := left - right
	factor := difference >> 63

	return left + factor*difference
}
