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
	Banner := os.Args[2]
	data, err := os.ReadFile(Banner + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	text := strings.Split(Arg[1], "\\n")
	slice := strings.Split(string(data), "\n")
	text = asciiart.IsNewline(text)
	 fmt.Print(asciiart.AsciiPrint(slice, text))
	
}
