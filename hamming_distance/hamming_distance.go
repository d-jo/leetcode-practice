package hamming_distance

import (
	"math"
)

func HammingDistance(x int, y int) int {

	diff := x ^ y
	count := 0

	for i := 1; i <= math.MaxInt32; i *= 2 {
		if diff&i == i {
			count++
		}
	}

	return count
}
