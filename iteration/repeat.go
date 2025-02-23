package iteration

import "strings"

func Repeat(character string) string {
	var repeat strings.Builder
	for i := 0; i < 5; i++ {
		repeat.WriteString(character)
	}
	return repeat.String()
}
