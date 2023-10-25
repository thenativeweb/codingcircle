package continuousmax

import (
	"errors"
)

func GetMax(values []int, k int) ([]int, error) {
	if len(values) == 0 {
		return []int{}, nil
	}
	if k <= 0 {
		return nil, errors.New("k must be greater than 0")
	}
	if k > len(values) {
		return nil, errors.New("k must be less than or equal to the length of values")
	}

	maxValues := make([]int, 0, len(values)-k+1)

	var window []int
	for i := 0; i < len(values); i++ {
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}

		for len(window) > 0 && values[window[len(window)-1]] <= values[i] {
			window = window[:len(window)-1]
		}

		window = append(window, i)

		if i >= k-1 {
			maxValues = append(maxValues, values[window[0]])
		}
	}

	return maxValues, nil
}
