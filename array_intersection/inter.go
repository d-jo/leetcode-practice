package inter

// use a hash map, for each element in A and B
// A adds 1 to map num
// B adds -1 to map num
// all the elements that are 0 in the map are the intersection
// wait doesnt work,
//
// a - 1,2,2,3
// b - 11,2,4,5
//
// this would result in no 2s in that solution
// but we need
func Intersect(nums1 []int, nums2 []int) []int {

	result := []int{}

	currA := make(map[int]int)
	currB := make(map[int]int)

	for _, v := range nums1 {
		currA[v]++
	}

	for _, v := range nums2 {
		currB[v]++
	}

	var longerTarget *map[int]int

	if len(currA) > len(currB) {
		longerTarget = &currA
	} else {
		longerTarget = &currB
	}

	for k, _ := range *longerTarget {
		aVal := currA[k]
		bVal := currB[k]

		if aVal > 0 && bVal > 0 {
			var shorterTarget *int
			if aVal == bVal {
				// add key to array aVal times
				shorterTarget = &aVal
			} else if aVal > bVal {
				// add bval
				shorterTarget = &bVal
			} else if aVal < bVal {
				// add aval
				shorterTarget = &aVal
			}

			for j := 0; j < *shorterTarget; j++ {
				result = append(result, k)
			}
		}
	}

	return result
}
