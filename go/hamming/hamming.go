package hamming

import (
	"errors"
)

//Distance compares the hamming distance between two equivalent length string representations of DNA
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return 0, errors.New("strands are of different length")
	}

	distance := 0
	for i, c := range ar {
		if c != br[i] {
			distance++
		}
	}

	return distance, nil
}
