func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var freq [26]int
	var window [26]int

	// create freq arr for s1
	for i, _:=range s1 {
		freq[s1[i]-'a']++
	}

	// first window
	for i, _ := range s2 {
		if i == len(s1) {
			break
		}
		window[s2[i]-'a']++
	}
	if window == freq {
		return true
	}

	// move the window now
	left:=0
	for right:=len(s1); right<len(s2);right++ {
		window[s2[right]-'a']++
		window[s2[left]-'a']--
		left++
		if window == freq {
			return true
		}
	}
	return false
}


