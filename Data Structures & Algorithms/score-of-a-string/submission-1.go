func scoreOfString(s string) int {
	sum:=0.0
	for i := 0; i < len(s)-1; i++ {
		sum += math.Abs(float64(s[i]) - float64(s[i+1]))
	}
	return int(sum)
}

// func abs(v int) int {
// 	if v < 0 {
// 	return -v
// 	}
// 	return v
// }