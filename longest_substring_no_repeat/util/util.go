package util

func HasRepeat(s string) bool {
	charMap := map[rune]bool{}

	for _, c := range s {
		if _, ok := charMap[c]; ok {
			return true
		}
		charMap[c] = true
	}
	return false
}

func RepeatArr(s string) int {
	placeMap := map[rune]bool{}
	longest := 0
	current := 0

	for _, c := range s {
		_, ok := placeMap[c]
		if ok {
			// we have seen it before
			if current > longest {
				longest = current
			}
			current = 1
			placeMap = map[rune]bool{}
		} else {
			// we haven't seen it before
			current++
		}
		placeMap[c] = true
	}

	if current > longest {
		longest = current
	}

	return longest
}

func LengthOfLongestSubstring(s string) int {
	return RepeatArr(s)
}
