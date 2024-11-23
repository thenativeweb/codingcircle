package parentheses

func IsBalanced(s string) bool {
	low, high := 0, 0
	for _, char := range s {
		switch char {
		case '(':
			low++
			high++
		case ')':
			low = max(low-1, 0)
			high--
		case '*':
			low = max(low-1, 0)
			high++
		}

		if high < 0 {
			return false
		}
	}

	return low == 0
}
