package reverse

import (
	"fmt"
	"strings"
)

func reverse_string(word string) string {
	// Firt convert the word to an array slice
	arrWord := strings.Split(word, "")

	if word == "" {
		return ""
	}

	left := 0
	right := len(word) - 1

	fmt.Println(arrWord[left], arrWord[right])

	for left < right {
		arrWord[left], arrWord[right] = arrWord[right], arrWord[left]

		left++
		right--
	}
	return strings.Join(arrWord, "")
}
