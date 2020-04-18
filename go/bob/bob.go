// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if len(remark) > 0 {
		if strings.ToLower(remark) == remark {
			if strings.LastIndex(remark, "?") == len(remark)-1 {
				return "Sure."
			}
		} else if strings.ToUpper(remark) == remark {
			if strings.Contains(remark, "?") {
				return "Calm down, I know what I'm doing!"
			}
			return "Whoa, chill out!"
		} else if strings.LastIndex(remark, "?") == len(remark)-1 {
			return "Sure."
		}
		return "Whatever."
	}
	return "Fine. Be that way!"
}
