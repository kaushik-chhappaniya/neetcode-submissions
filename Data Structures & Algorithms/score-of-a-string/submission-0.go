func scoreOfString(s string) int {
	var sum int = 0
	for i := 0; i < len(s)-1; i++ {
		sum += abs(int(s[i]) - int(s[i+1]))
	}
	return sum
}

func abs(v int) int {
	if v < 0 {
	return -v
	}
	return v
}