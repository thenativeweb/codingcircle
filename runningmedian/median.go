package runningmedian

import "sort"

func GetMedians(sequence []int) []float64 {
	medians := make([]float64, len(sequence))

	tempSequence := []int{}
	for i, element := range sequence {
		tempSequence = append(tempSequence, element)
		sort.Ints(tempSequence)

		var median float64
		n := len(tempSequence)
		if n%2 == 1 {
			median = float64(tempSequence[n/2])
		} else {
			median = float64(tempSequence[n/2-1]+tempSequence[n/2]) / 2
		}

		medians[i] = median
	}

	return medians
}
