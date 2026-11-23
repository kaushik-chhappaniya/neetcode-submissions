func countConsistentStrings(allowed string, words []string) int {
	// create map of boolean to track which letters are there
	hm:=make(map[byte]bool)
	for i := 0; i < len(allowed); i++ {
		hm[allowed[i]] = true
	}
	count:=len(words)
	// use nested loop to iterate over words and their letters
	// if some unknown letter is found break loop and reduce count. 
	for _, word := range words {
		for i:=0; i< len(word); i++ {
			if !hm[word[i]] {
				count--
				break
			} 
		}
	}
	return count
}