import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums[:])
	var result [][]int
	
	for i,v:= range nums {
		if v > 0 {
			break
		}
		if i > 0 && v == nums[i-1] {
			continue
		}
		
		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := v + nums[l] +nums[r]
			if threeSum > 0{
				r--
			} else if threeSum < 0 {
				l++
			} else {
				result = append(result, []int{v, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return result
}

