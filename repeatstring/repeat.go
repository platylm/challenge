package repeatstring

import (
	"strings"
)

func repeatedString(s string, n int) int {
	count := strings.Count(s, "a")

	repeat := n / len(s)

	splitString := strings.Split(s, "")

	portion := n % len(s)
	slicePortion := splitString[:portion]
	convertSliceToString := strings.Join(slicePortion, "")

	totalA := (count * repeat) + strings.Count(convertSliceToString, "a")

	return totalA
}
