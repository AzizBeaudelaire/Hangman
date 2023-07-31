package hangman

import (
	"fmt"
	"strings"
)

// TestFile , tests if the file contain false words or empty lines
func TestFile(file []string) bool {
	for _, word := range file {
		if len(word) > 0 {
			for _, letter := range strings.ToUpper(word) {
				if letter < 65 || letter > 90 {
					fmt.Println("\033[31m", "This file must contain only letters. Please modify the words.txt file.")
					return false
				} else {
					continue
				}
			}
			continue
		} else {
			fmt.Println("\033[31m", "The words.txt file contain empty lines. Please modify it.")
			return false
		}
	}
	return true
}
