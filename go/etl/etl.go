package etl

import "strings"

//Transform takes in a map of scoring and converts it in to an old scoring system
func Transform(oldScores map[int][]string) map[string]int {

	scores := make(map[string]int)

	for key, elements := range oldScores {
		for _, element := range elements {
			scores[strings.ToLower(string(element))] = key
		}
	}

	return scores
}
