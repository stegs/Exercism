package strand

//ToRNA takes a dna string and returns the rna string
func ToRNA(dna string) string {
	rnaMap := make(map[string]string)
	rnaMap["G"] = "C"
	rnaMap["C"] = "G"
	rnaMap["T"] = "A"
	rnaMap["A"] = "U"

	rna := ""

	for _, c := range dna {
		rna += rnaMap[string(c)]
	}

	return rna
}
