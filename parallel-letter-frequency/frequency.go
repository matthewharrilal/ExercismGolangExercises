package letter

import "fmt"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{} // Instantiating the map to contain the frequency
	for _, r := range s {
		m[r]++ // Increment to the value of occurences
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	// Concurrently calculates the frequency of occurences of letters in a string
	channel := make(chan FreqMap) // Going to be writing the frequency of occurences of the letters through the channels
	result := FreqMap{}
	for _, word := range texts {
		// Once we are able to calculate the value of the letter we are currently on we want to be

		go func(word string) {
			//fmt.Println(Frequency(word))
			channel <- Frequency(word) // Write histogram to channel for each word
		}(word) // Execute anonymous func
	}

	// SO what do we know, that texts is an array of words

	for range texts { // Iterate over the words in the texts
		//fmt.Println(texts)
		for val := range <-channel {
			fmt.Println(val)
			//result[letter] ++
		}
	}

	fmt.Println(result)
	return result
}
