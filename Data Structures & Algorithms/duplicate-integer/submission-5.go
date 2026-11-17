func hasDuplicate(nums []int) bool {
    hm := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if !hm[nums[i]] {
			hm[nums[i]] = true
		} else {
			return true
		}
	}
    return false
}
