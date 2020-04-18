package scale

//Scale is a function that creates the full musical scale based on the initial tonic and the provided interval
func Scale(tonic string, interval string) []string {
	var chromatic = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	var chromaticFlat = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "Fb", "F", "Gb", "G", "Ab"}
	var activeBaseScale = make([]string, 12)

	if 
}
