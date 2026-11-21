func findLucky(arr []int) int {
	hm := make(map[int]int)
	// Create hashmap for storing freq
	for _,v := range arr {
		hm[v]++
	}
	res := -1
	// iterate over hashmap and check if num and freq is same, keep a max marker to maintain large number
	for num, freq := range hm {
		if num == freq {
			if num > res {
				res = num
			}
		}
	}
	return res
}
