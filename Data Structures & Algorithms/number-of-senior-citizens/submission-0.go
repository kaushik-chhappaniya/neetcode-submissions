func countSeniors(details []string) int {
	count:=0
    for _,s:= range details {
		if v, _ := strconv.Atoi(s[11:13]); v > 60 {
			count++
	}
	}
	return count
}