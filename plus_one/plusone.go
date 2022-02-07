package plusone

import "fmt"

// digits is msb, msb-1
// ex: 1034 -> [1, 0, 3, 4]
func PlusOne(digits []int) []int {

	var carry int = 0
	result := make([]int, 0)
	modify := 1

	for i := len(digits) - 1; i >= 0; i-- {
		tmpSum := digits[i] + carry + modify
		digit := tmpSum % 10

		fmt.Println(i, tmpSum, digit)

		if tmpSum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		result = append([]int{digit}, result...)
		modify = 0
	}

	if carry == 1 {
		result = append([]int{1}, result...)
	}

	return result
}
