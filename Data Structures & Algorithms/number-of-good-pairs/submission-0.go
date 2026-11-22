func numIdenticalPairs(nums []int) int {
    count := 0
	hm := make(map[int]int)
	for _,i := range nums {
		count += hm[i]
		hm[i]++
	}
	return count
}