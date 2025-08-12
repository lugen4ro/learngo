package iteration

import "strings"

func Repeat(character string, n int) string {
	var repeated strings.Builder
	// for i := 0; i < n; i++ {
	for range n {
		repeated.WriteString(character)
	}
	return repeated.String()
}
