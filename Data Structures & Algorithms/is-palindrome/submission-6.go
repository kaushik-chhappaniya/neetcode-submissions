func isPalindrome(s string) bool {
	// reg := regexp.MustCompile("[0-9 !@#$%^:;-_=~&*(){}|\\,?>'`<.?\\/]+")
	// k := strings.ToLower(reg.ReplaceAllString(s, ""))
	// if len(k) <= 1 {
	// 	return true
	// }
	cleanStr := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r)  || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, s)
	cleanStr = strings.ToLower(cleanStr)
	if len(cleanStr) == 0 { 
		return true
	}

	for i,j:=0,len(cleanStr)-1; i<=j; i,j=i+1,j-1{
		if cleanStr[i] != cleanStr[j] {
			return false
		}
	}
	return true
}

// func removeAlphanumeric(s string) string {
// 	return strings.Map(func(r rune) rune {
// 		// If the character is a letter or a digit, drop it (-1)
// 		if unicode.IsLetter(r) || unicode.IsDigit(r) {
// 			return -1
// 		}
// 		// Keep special characters, spaces, and punctuation
// 		return r
// 	}, s)
// }
