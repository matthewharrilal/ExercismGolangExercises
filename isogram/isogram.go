package isogram

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	frequencyCount := make(map[string]int) // Containing the frequency count of the letters in the world
	word = strings.ToLower(word)
	for _, letter := range word {
		// Ignore spaces and hyphens
		if letter != ' ' && letter != '-' { // Using single quotes to denote character as opposed to string
			_, ok := frequencyCount[string(letter)]

			if ok { // Meaning callback of operation returned true
				return false
			} else {
				frequencyCount[string(letter)]++
			}
		}
	}
	fmt.Println(frequencyCount)
	return true
}
