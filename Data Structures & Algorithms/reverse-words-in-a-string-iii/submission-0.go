func reverseWords(s string) string {
	char := []byte(s)
	start := 0
	for end := 0; end <= len(char); end++ {
		if end == len(char) || char[end] == ' ' {
			for l, r := start, end-1; l < r; l, r = l+1, r-1 {
				char[l], char[r] = char[r], char[l]
			}
			start = end + 1
		}
	}
	return string(char)
}
