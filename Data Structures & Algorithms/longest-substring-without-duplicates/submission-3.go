func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var l,res int = 0,0
	hm := make(map[byte]int)
	for r:=0; r< len(s) ; r++ {
		if idx, ok := hm[s[r]]; ok {
			l = max(idx+1, l)
		}
		hm[s[r]] = r
		if r-l+1 >res {
			res = r-l+1
		}
	}
	return res
}