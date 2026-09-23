func getConcatenation(nums []int) []int {
	n:=len(nums)
    ans := make([]int, n*2)
	copy(ans,nums)
	for i:=n; i<len(ans); i++ {
		ans[i] = nums[i-n]
	}
	return ans
}
