func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var window [26]int
	var freq [26]int
	// Create freq arr of s1 small string
	for i:=0;i<len(s1);i++{
		freq[s1[i]-'a']++
	}

	

	// Create first sliding window from s2
	for i:=0; i< len(s1);i++ {
		window[s2[i]-'a']++
	}
	// First check the first window and freq arr
	if freq == window {
		return true
	}



	// now vary the sliding windows
	left := 0
	for right:=len(s1);right<len(s2);right++{
		// Add new char from left side
		window[s2[right]-'a']++
		// Remoev the first char from window
		window[s2[left]-'a']--
		left++
		if freq == window {
			return true
		}
	}
		return false
}


