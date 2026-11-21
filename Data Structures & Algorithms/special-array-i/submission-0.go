func isArraySpecial(nums []int) bool {
	// iterate over array and do bitwise and that will give odd or even
    for i := 1; i < len(nums); i++ {
		if (nums[i-1] & 1) == (nums[i] & 1) {
			return false
		}
	}
	return true
}