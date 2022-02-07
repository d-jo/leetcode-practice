package reverse_string

// must be done in place
func ReverseString(s []byte) {

	var j int = 0
	for i := 0; ; i++ {
		j = len(s) - 1 - i
		if i == j || j < i {
			return
		}

		// swap
		s[i], s[j] = s[j], s[i]
	}

}
