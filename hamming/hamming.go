package hamming

import ("fmt"
		"errors")

func Distance(a, b string) (int, error) {
	hammingDistance := 0 // Represent counter for errors
	fmt.Println(len(a), len(b))

	if len(a) != len(b) {
		return 0, errors.New("Of different lengths")
	}
	for index, letter := range a {

		if letter != rune(b[index]) {
			hammingDistance++
		}

	}

	return hammingDistance, nil
}
