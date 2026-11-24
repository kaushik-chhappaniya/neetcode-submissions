func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	var arr [26]int
	for _, v:= range magazine {
		// increase counter for each char in magazine
		arr[v - 'a']++
	}
	for _, v := range ransomNote  {
		// remove each char from arr that is in ransomNote 
		arr[v - 'a']--
		// check if the array has negative values for the char, if yes meaning magazine doesn't have enough characters required from ransomNote
		if arr[v - 'a'] < 0 {
			return false
		}
	}
	return true
}
