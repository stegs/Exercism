package raindrops

import "strconv"

//Convert is a function that takes an integer input and outputs the raindrop representation.
//If no representation is available it returns the string version of the input
func Convert(input int) string {
	output := ""

	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}

	if len(output) == 0 {
		return strconv.Itoa(input)
	}

	return output
}
