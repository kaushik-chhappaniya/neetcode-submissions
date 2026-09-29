func findMaxConsecutiveOnes(nums []int) int {
	count, maxCount := 0,0
	for i,v:=range nums {
		if v != 0 {
			count++
		} 
		if v == 0 || i == len(nums)-1 {
			maxCount = max(maxCount, count)
			count = 0 
		}
	}
	return maxCount
}
