func characterReplacement(s string, k int) int {
	hm := make(map[byte]int)
	l,maxf, res := 0,0,0
	for i:=0; i< len(s); i++ {
		hm[s[i]]++
		if hm[s[i]] > maxf {
			maxf = hm[s[i]]
		}
		for (i - l + 1) - maxf > k {
			hm[s[l]]--
			l++
		}

		if i - l + 1 > res {
			res = i - l + 1
		}
	}
	return res
}
