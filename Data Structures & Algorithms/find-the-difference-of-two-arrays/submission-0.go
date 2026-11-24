func findDifference(nums1 []int, nums2 []int) [][]int {
	// create hashmaps for 2 arrays
	hm1 := make(map[int]struct{})
	hm2 := make(map[int]struct{})
	
	for _, num := range nums1 {
		hm1[num] = struct{}{}
	}
	
	for _, num := range nums2 {
		hm2[num] = struct{}{}
	}
	// Create 2 arrays for results
	res1, res2 := []int{}, []int{}
	for num := range hm1 {
		if _, ok := hm2[num]; !ok {
			res1 = append(res1, num)
		}
	}
	for num := range hm2 {
		if _, ok := hm1[num]; !ok {
			res2 = append(res2, num)
		}
	}

	return [][]int{res1, res2}
}
