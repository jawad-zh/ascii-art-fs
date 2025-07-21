package asciiartoutput

// ValidCharacter checks if all characters in the input string are valid ASCII characters (32 to 126)
func ValidCharacter(input string) bool {
	for _, char := range input {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
