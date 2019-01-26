package diffsquares

import (
	"fmt"
	"math"
)

func DifferenceOfSquares(numSlice []int) int {
	sum := 0
	squaredSum := 0.
	for _, val := range numSlice {
		sum += val
	}

	for _, val := range numSlice {
		squaredSum += math.Pow(float64(val), 2)
	}
	fmt.Println(sum)
	exponentialSum := math.Pow(float64(sum), 2)

	return int(exponentialSum) - int(squaredSum)
}
