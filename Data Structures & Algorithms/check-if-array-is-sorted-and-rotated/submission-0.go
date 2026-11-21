func check(nums []int) bool {
    l, r := 0, len(nums)
	for i := 0; i < r; i++ {
		if nums[i] > nums[(i+1)%r] {
			l++
			if l > 1 {
				return false
			}
		}
	}
	return true
}