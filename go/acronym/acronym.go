// Package acronym contains methods to create acronyms from words
package acronym

import (
	"strings"
)

// Abbreviate returns the acronym of the provided string
func Abbreviate(s string) string {
	replacer := strings.NewReplacer("-", " ", "_", " ") //Does not account for other potential punctuation!

	sArray := strings.Split(replacer.Replace(s), " ")

	abr := ""

	for i := 0; i < len(sArray); i++ {
		if len(sArray[i]) > 0 {
			abr += strings.ToUpper(sArray[i][0:1])
		}
	}

	return abr
}
