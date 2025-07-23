package asciiartoutput

import (
	"fmt"
)

// check if the input is valid and the charachters is in the ascii range
func ValidateInput(Arg []string) bool {
	if len(Arg) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER] \nEX: go run .  something standard")
		return false
	}

	if Arg[2] != "standard" && Arg[2] != "shadow" && Arg[2] != "thinkertoy" {
		fmt.Println("The Banner must be one of this list : \n1.standard\n2.shadow\n3.thinkertoy")
		return false
	}

	if !ValidCharacter(Arg[1]) {
		fmt.Println("Invalid characters in input. Only printable ASCII characters are allowed.")
		return false
	} else {
		return true
	}
}
