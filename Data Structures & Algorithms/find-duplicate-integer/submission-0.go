func findDuplicate(nums []int) int {
    for _, num := range(nums) {
		id := abs(num) - 1
		if nums[id] < 0 {
			return abs(num)
		}
		nums[id] *= -1
	}
	return -1
}

func abs(num int ) int {
	if num < 0 {
		return -num
	}
	return num
}
