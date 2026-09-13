func findMin(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	minNum := math.MaxInt
	// for i,j := 0,len(nums)-1; i < j; i,j =i+1, j-1 {
	// 	minNum = min(minNum, nums[i], nums[j])
	// }
	for i := 0; i< len(nums); i++ {
		minNum = min(minNum, nums[i])
	}
	return minNum
}
