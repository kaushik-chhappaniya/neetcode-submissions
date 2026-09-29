func lengthOfLastWord(s string) int {
	if len(s) == 1 {
		return 1
	}
	count := 0
	for i := len(s)-1; i >= 1; i-- {
		if string(s[i]) != " " {
			count++
			if string(s[i-1]) == " " {
				return count
			}
		}
	}
		return count
	// str:= strings.TrimSpace(s)
	// start,end := 0,0
	// n:=len(str)
	// for i:=0; i< n; i++ {
	// 	if str[i] == ' '{
	// 		start = i
	// 		end = i
	// 	} else {
	// 		end++
	// 	}
	// }
	// return end -start
}
