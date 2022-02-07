package strstr

func StrStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

outter:
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] {
			for j := 0; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					continue outter
				}
			}

			return i
		}
	}

	return -1
}
