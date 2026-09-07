func countBits(n int) []int {
	var arr []int

	for i := range n+1 {
		count := 0
		for i!= 0 {
			i &= i-1
			count++
		}
		arr = append(arr, count)
	}
	return arr
}
