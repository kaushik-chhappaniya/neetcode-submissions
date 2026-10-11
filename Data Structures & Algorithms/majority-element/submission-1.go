func majorityElement(nums []int) int {
    hm := make(map[int]int)
	for _,v := range nums {
		hm[v]++
		if val, ok := hm[v]; ok && val > len(nums)/2{
			return v
		}
	}
	return -1
}
