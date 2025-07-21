package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiartoutput/functions"
)

func main() {
	Arg := os.Args
	if !asciiart.ValidateInput(Arg) {
		return
	}
	Banner := os.Args[3]
	data, err := os.ReadFile(Banner + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	text := strings.Split(Arg[2], "\\n")
	slice := strings.Split(string(data), "\n")
	text = asciiart.IsNewline(text)
	output := asciiart.AsciiPrint(slice, text)
	err = os.WriteFile(strings.TrimPrefix(os.Args[1], "--output="), []byte(output), 0o644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
