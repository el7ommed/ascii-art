package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	//welcome to the main function, its only job is to send the text
	//from terminal to the ArtGenerator function and print it to console
	//actual code starts at line #41 (...ish)
	//please note that you need to use double quotes (") to provide multiple words seperated by a space
	//don't blame me, I'm not the one who designed the terminal ¯\_(ツ)_/¯
	if len(os.Args) < 2 {
		fmt.Println("Please provide text to generate ASCII Art!\nAny extra arguments will be ignored")
		return
	}
	if os.Args[1] == "" {
		return
	}
	array := strings.Split(os.Args[1], "\\n")
	for i := 0; i < len(array); i++ {
		if i+1 < len(array) {
			if array[i] != "" {
				fmt.Println(ArtGenerator(array[i]))
			} else {
				fmt.Println()
			}
		} else if i+1 == len(array) { //the last string in the slice is special
			if array[i] != "" {
				fmt.Println(ArtGenerator(array[i]))
			} else if array[i] == "" && array[i-1] != "" {
				fmt.Println()
			}
		}
	}
}

func ArtGenerator(input string) string {
	if len(input) == 0 {
		return ""
	}
	output := make([]string, 8)
	content, _ := os.ReadFile("standard.txt")
	font := strings.Split(string(content), "\n") //this is the font file

	for _, runeValue := range input {
		//this is where the function checks the letters one by one and add them
		//to the output string as large ascii letters within 8 lines
		//valid characters are #32 (space) to #126 tilde(~) in the ascii table
		//it will replace any invalid characters with a #63 question mark(?)
		var letterIdx int
		if utf8.RuneLen(runeValue) == 1 && runeValue >= 32 && runeValue <= 126 {
			letterIdx = int(runeValue) - 32
		} else {
			letterIdx = 63 - 32
		}

		for j := 0; j < 8; j++ {
			output[j] += font[(letterIdx*9)+j+1]
		}
	}

	// for i := 0; i < len(input); i++ {
	// 	if byte(input[i]) >= 32 && byte(input[i]) <= 126 {
	// 		for j := 0; j < 8; j++ {
	// 			var letterIdx int = int(input[i]) - 32
	// 			output[j] += font[(letterIdx*9)+j]
	// 		}
	// 	} else {
	// 		for j := 0; j < 8; j++ {
	// 			var letterIdx int = 63 - 32
	// 			output[j] += font[(letterIdx*9)+j]
	// 		}
	// 	}
	// }

	//return it as a string so it can be printed by the main function
	return strings.Join(output, "\n")
}
