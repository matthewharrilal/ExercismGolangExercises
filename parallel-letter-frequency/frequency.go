package letter

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

	// Instantiate channel that we are going to be passing maps through
	channel := make(chan FreqMap) // Going to be writing the frequency of occurences of the letters through the channels

	// Store result of overall text in result map
	result := FreqMap{}

	// Iterate over words in the given text
	for _, word := range texts {

		// Spin a separate go routine for every word to calculate the frequency for the word
		go func(word string) {
			channel <- Frequency(word) // Write histogram to channel for each word
		}(word) // Execute anonymous func
	}

	// Iterate over the words given to recieve all words written through the channel or else sender will block .... no one to recieve
	for range texts { // Recieve from channel from number of times you send through the channel
		for letter, frequency := range <-channel {
			// For the word written through the channel recieve and calculate frequency in result map ... will only get the first word if not iterating over the range of the texts
			result[letter] += frequency
		}
	}
	return result
}
