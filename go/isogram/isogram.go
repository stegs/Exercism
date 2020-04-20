package isogram

import (
	"strings"
)

//IsIsogram takes a string input and returns a boolean value saying whether the string is an isogram
func IsIsogram(input string) bool {

	words := strings.Split(strings.ToLower(input), " ")

	for _, w := range words {
		for _, c := range w {
			if c == '-' || c == ' ' {
				continue
			}
			if strings.Count(w, string(c)) > 1 {
				return false
			}
		}
	}

	return true
}
