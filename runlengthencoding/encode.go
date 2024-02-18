package runlengthencoding

import (
	"strconv"
	"strings"
)

func Encode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder

	count := 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(input[i-1])
			count = 1
		}
	}

	result.WriteString(strconv.Itoa(count))
	result.WriteByte(input[len(input)-1])

	return result.String()
}
