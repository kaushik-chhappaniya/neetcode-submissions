func characterReplacement(s string, k int) int {
	res := 0
	charS := make(map[byte]bool)

	for i:=0; i< len(s); i++ {
		charS[s[i]] = true
	}

	for c := range charS {
		count, l := 0,0
		for r := 0; r < len(s); r++ {
			if s[r] == c {
				count++
			}
			for (r - l + 1) - count > k {
				if s[l] == c {
					count--
				}
				l++
			}
			res = max(res, r - l + 1)
		}
	}
	return res
}
