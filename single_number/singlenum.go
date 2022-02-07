package singlenum

// solution should be linear in time and constant in space
// how to do constant in spacce? this approach is linear
func SingleNumberWrong(nums []int) int {
	vals := make(map[int]int)
	for i, v := range nums {
		vals[v]++
		if vals[v] >= 2 {
			return i
		}
	}

	return -1
}

// I googled this... should have figured it out
func SingleNumber(nums []int) int {
	curr := int(0)

	for _, v := range nums {
		curr ^= v
	}

	return curr
}
