package collatzconjecture

import "errors"

//CollatzConjecture calculates the number of calculations it takes to get from the input to the number one.
//It returns the number of steps and the error.
func CollatzConjecture(input int) (int, error) {
	numSteps := 0

	if input < 1 {
		return 0, errors.New("the number must be 1 or greater")
	}

	for {
		if input == 1 {
			break
		} else if input%2 == 0 {
			input = input / 2
		} else {
			input = (3 * input) + 1
		}
		numSteps++
	}

	return numSteps, nil
}
