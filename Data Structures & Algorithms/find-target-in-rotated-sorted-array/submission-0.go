func search(nums []int, target int) int {
	l , r:= 0, len(nums)-1
	for l <= r {
		mid := (l + r) /2
		// check if targ is at mid
		if nums[mid] == target {
			return mid
		}
		// check if left less than mid - means pivot is on the right 3 4 5 6 0 1 2
		if nums[l] <= nums[mid] {
			// target lies in the right part
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else { // means pivot is on the left side 5 6 0 1 2 3 4
			// target lies in the left part
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
