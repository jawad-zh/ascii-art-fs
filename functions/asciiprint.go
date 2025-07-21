package asciiartoutput

// AsciiPrint prints the ASCII art representation of the input strings
func AsciiPrint(slice, texts []string) string {
	var output string
	for _, text := range texts {
		if text == "" {
			output += "\n"
			continue
		}
		for j := 1; j < 9; j++ {
			for _, char := range text {
				output += slice[(int(char)-32)*9+j]
			}
			output += "\n"
		}

	}
	return output
}
