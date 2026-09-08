func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if target > nums[len(nums)-1] || target < nums[0] {
		return -1
	}
	l,r := 0,len(nums)-1
	for l<r {
		mid := l + (r-l)/ 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1 
			}
		}
		if l < len(nums) && nums[l] == target {
			return l
		}
		return -1
}
