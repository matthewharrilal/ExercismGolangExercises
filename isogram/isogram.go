package isogram

func DetermineIsogram(word string) bool {
	frequencyCount := make(map[string]int) // Containing the frequency count of the letters in the world

	for _, letter := range word {
		// Ignore spaces and hyphens
		if letter != ' ' && letter == '-' { // Using single quotes to denote character as opposed to string
			_, ok := frequencyCount[string(letter)]

			if ok { // Meaning callback of operation returned true
				return false
			} else {
				frequencyCount[string(letter)] ++ 
			}
		}
	}
	return true
}