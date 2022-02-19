package power_of_three

import (
	"math"
)

/*
	log_3(n^3) = 3 * log_3(n)
	3 * log_3(n) = 3 * (log(n) / log(3)) = res

	if res is an integer, its a power of 3
*/
func IsPowerOfThree(n int) bool {
	// how this works
	res := 3.0 * (math.Log(float64(n)) / math.Log(3.0))
	val := math.Abs(res - math.Trunc(res))
	return val <= 0.000000000000002
}
