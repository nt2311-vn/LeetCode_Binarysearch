package smallestgreater

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)
	for l < r {
		m := l + (r-l)/2
		if letters[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}

	if l == len(letters) {
		return letters[0]
	}

	return letters[l]
}
