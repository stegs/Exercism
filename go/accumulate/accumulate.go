package accumulate

//Accumulate takes a string array and a converter function and returns an array of converted strings
func Accumulate(given []string, converter func(string) string) []string {

	accumulated := make([]string, len(given))

	for i, w := range given {
		accumulated[i] = converter(w)
	}

	return accumulated
}
