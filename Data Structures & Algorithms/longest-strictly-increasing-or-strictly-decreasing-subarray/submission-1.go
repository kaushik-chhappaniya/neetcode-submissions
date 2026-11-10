func longestMonotonicSubarray(nums []int) int {
	n := len(nums)
	res := 1

	for i := 0; i < n-1; i++ {
		currLen := 1
		for j := i+1; j < n; j++ {
			if nums[j] == nums[j-1] || ((nums[i] < nums[i+1]) != (nums[j-1] < nums[j])) {
				break
			}
			currLen++
		}
		if currLen > res {
			res = currLen
		}
	}
	return res
}