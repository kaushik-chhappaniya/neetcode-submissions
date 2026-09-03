func minWindow(s string, t string) string {
    if len(t) > len(s) {
		return ""
	}
	var need [128]int
	var window [128]int
	// Create freq arr of t 
	for i:=0; i< len(t);i++ {
		need[t[i]]++
	}
	required := 0
	// find out unique chars in t
	for i:=0 ;i<128;i++{
		if need[i]>0 {
			required++
		}
	}

	have :=0
	left := 0

	bestStart:=0
	bestLen := len(s)+1
	for right:=0; right<len(s);right++ {
		c:=s[right]
		window[c]++
		// this char requirement is satisfied here
		if need[c] > 0 && window[c] == need[c] {
			have++
		}

		// current window is valid try to make it smaller now
		for have == required {
			currentLen := right - left +1
			if currentLen < bestLen {
				bestLen = currentLen
				bestStart = left
			}
			leftChar := s[left]
			window[leftChar]--

			// Removing this left char broke a requirement check
			if need[leftChar] > 0 && window[leftChar] <need[leftChar] {
				have--
			}
			left++
		}
	}
	if bestLen == len(s)+1 {
		return ""
	}
	return s[bestStart:bestStart+bestLen]

}
