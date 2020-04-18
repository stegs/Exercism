package scrabble

import "strings"

//Score takes an input and outputs the scrabble score
func Score(input string) int {
	values := make(map[string]int)

	values["a"], values["e"], values["i"], values["o"], values["u"], values["l"], values["n"], values["r"], values["s"], values["t"] = 1, 1, 1, 1, 1, 1, 1, 1, 1, 1
	values["d"], values["g"] = 2, 2
	values["b"], values["c"], values["m"], values["p"] = 3, 3, 3, 3
	values["f"], values["h"], values["v"], values["w"], values["y"] = 4, 4, 4, 4, 4
	values["k"] = 5
	values["j"], values["x"] = 8, 8
	values["q"], values["z"] = 10, 10

	total := 0

	for _, c := range input {
		total += values[strings.ToLower(string(c))]
	}

	return total
}
