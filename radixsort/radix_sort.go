package radixsort

import (
	"math"
)

func max(values []int) int {
	maxValue := math.MinInt
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func countDigits(number int) int {
	if number == 0 {
		return 1
	}
	return int(math.Log10(float64(number))) + 1
}

func getDigit(number, n int) int {
	return (number / int(math.Pow(10, float64(n)))) % 10
}

func Sort(values []int) []int {
	digits := countDigits(max(values))

	for k := 0; k < digits; k++ {
		buckets := make([][]int, 10)

		for _, value := range values {
			digit := getDigit(value, k)
			buckets[digit] = append(buckets[digit], value)
		}

		values = []int{}
		for _, bucket := range buckets {
			values = append(values, bucket...)
		}
	}

	return values
}
