// Package twofer contains the ShareWith function.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith returns the string "One for 'x', one for me." based on the input from an external function.
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	}

	return "One for you, one for me."
}
