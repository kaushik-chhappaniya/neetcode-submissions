func appendCharacters(s string, t string) int {
    // check if its substr or not if not then app
	// check from max len
	// reduce 1 1
	j:=0
	for i:=0; i< len(s); i++ {
		if j<len(t) && s[i] == t[j] {
			j++
		}
		if j == len(t) {
			return 0
		}
	}
	return len(t)-j
}