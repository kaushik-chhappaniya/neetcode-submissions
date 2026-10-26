import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var res [][]int
	for i,v := range nums {
		if v > 0 {
			break}
		if i > 0 && v== nums[i-1] {
			continue
		}
		l,r  := i+1, len(nums)-1
		for l < r {
			sum := v + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0{
				r--
			} else {
				res = append(res, []int{nums[l],nums[r],v})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		} 
	 }
	 return res
}