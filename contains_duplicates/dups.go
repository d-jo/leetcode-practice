package dups

func ContainsDuplicate(nums []int) bool {
	items := make(map[int]bool)

	for _, i := range nums {
		if _, ok := items[i]; ok {
			return true
		}
		items[i] = true
	}
	return false
}
