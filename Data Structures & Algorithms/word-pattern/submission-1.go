func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s) // split by blank spaces
	// if lengths of both pattern and string are unequal meaning they're already not matching
	if len(words) != len(pattern) { 
		return false
	}

	// make map for word to pattern matching
    wordToPattern := make(map[byte]int)
	// lengths same so iterate over any one length
	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		w := words[i]
		if id, ok := wordToPattern[c]; ok {
			if words[id] != w {
				return false
			}
		} else {
			for _, val := range wordToPattern {
				if words[val] == w {
					return false
				}
			}
			wordToPattern[c] = i
		}
	}
	return true
}
