func removeDuplicates(nums []int) int {
	wr :=0 
	for i:=0; i< len(nums);i++ {
		if nums[wr] != nums[i] {
			wr++
			nums[wr] = nums[i]
		}
	}
	return wr + 1
}
