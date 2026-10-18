func stringMatching(words []string) []string {
    kmp := func(word1, word2 string) int {
		lps := make([]int, len(word2))
		prevLPS, i := 0,1

		for i < len(word2) {
			if word2[i] == word2[prevLPS] {
				lps[i] = prevLPS + 1
				prevLPS++
				i++
			} else if prevLPS == 0 {
				lps[i] = 0
				i++
			} else {
				prevLPS = lps[prevLPS - 1]
			}
		}
		i,j := 0,0
		for i < len(word1) {
			if word1[i] == word2[j] {
				i++
				j++
			} else {
				if j == 0 {
					i++
				} else {
					j = lps[j-1]
				}
			}
			if j == len(word2) {
				return i - len(word2)
			}
		}
		return -1
	}
	res := []string{}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for i :=0; i < len(words); i++ {
		for j := i+1; j< len(words); j++ {
			if kmp(words[j], words[i]) != -1 {
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}	