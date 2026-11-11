func pivotIndex(nums []int) int {
	total := 0
	// compute total sum first
	for _, num := range nums {
		total += num
	}

	left :=0
	for i :=0;i< len(nums);i++ {
		//calculate right sum at the same time, and match with leftsum. 
		right:=total-left -nums[i]
		if left == right {
			return i
		}
		left += nums[i]
	}
	return -1
}
