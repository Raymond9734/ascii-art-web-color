package Ascii

import (
	"fmt"
	"strings"
)

// PrintBanner prints the input string using the loaded banner characters.
func PrintBanner(line, fileName string) (string, error) {
	outPut := make([][]string, 8) // Output slice to store the banner lines.

	// Checks if the banner file size is not altered.
	filePath, err := FileCheck(fileName)
	if err != nil {
		return "", err
	}

	banner, err := LoadBanner(filePath) // Load the banner characters
	if err != nil {
		return "", err
	}
	for _, char := range line {
		if char < 32 || char > 126 {

			return "", fmt.Errorf("character out of range: %q", char)
		}
		if ascii, Ok := banner[char]; Ok {

			// If the character is found, split it into lines and append to the output
			asciiLines := strings.Split(ascii, "\n")
			for i := 0; i < len(asciiLines); i++ {
				outPut[i] = append(outPut[i], asciiLines[i])
			}
		} else {
			// If the character is not found, print an error message and continue
			fmt.Printf("Charachter not found: %q\n", char)
			continue
		}
	}
	var outPutAscii string = "\n"
	// Print the assembled output lines
	for _, line := range outPut {
		outPutAscii += strings.Join(line, "") + "\n"
	}
	return outPutAscii, nil
}
