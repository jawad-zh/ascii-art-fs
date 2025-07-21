package asciiartoutput

// Check if the input is empty or contains only newlines
func IsNewline(input []string) []string {
	for index, text := range input {
		if text != "" {
			break
		}
		if index == len(input)-1 && text == "" {
			return input[:len(input)-1]
		}
	}
	return input
}
