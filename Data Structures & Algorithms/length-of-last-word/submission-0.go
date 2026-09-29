func lengthOfLastWord(s string) int {
	str:= strings.TrimSpace(s)
	start,end := 0,0
	n:=len(str)
	for i:=0; i< n; i++ {
		if str[i] == ' '{
			start = i
			end = i
		} else {
			end++
		}
	}
	return end -start
}
