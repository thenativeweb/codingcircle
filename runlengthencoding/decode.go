package runlengthencoding

import (
	"strconv"
	"strings"
	"unicode"
)

func Decode(input string) string {
	var result strings.Builder
	var countStr strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			countStr.WriteRune(char)
		} else {
			count, _ := strconv.Atoi(countStr.String())
			result.WriteString(strings.Repeat(string(char), count))
			countStr.Reset()
		}
	}

	return result.String()
}
