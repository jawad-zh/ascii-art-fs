package asciiartoutput

import (
	"fmt"
	"strings"
)

// check if the input is valid and the charachters is in the ascii range
func ValidateInput(Arg []string) bool {
	count:=Count()
	if count==3{
		if !strings.HasPrefix(Arg[1], "--output=") || !strings.HasSuffix(Arg[1], ".txt") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> something standard")
			return false
		}
		if Arg[3] != "standard" && Arg[3] != "shadow" && Arg[3] != "thinkertoy" {
		fmt.Println("The Banner must be one of this list : \n1.standard\n2.shadow\n3.thinkertoy")
		return false
	}
	if !ValidCharacter(Arg[2]) {
		fmt.Println("Invalid characters in input. Only printable ASCII characters are allowed.")
		return false
	} 
	}else if count==2{
		if Arg[2] != "standard" && Arg[2] != "shadow" && Arg[2] != "thinkertoy" {
		fmt.Println("The Banner must be one of this list : \n1.standard\n2.shadow\n3.thinkertoy")
		}
	if !ValidCharacter(Arg[1]) {
		fmt.Println("Invalid characters in input. Only printable ASCII characters are allowed.")
		return false
		} 
	}
	return true
	}