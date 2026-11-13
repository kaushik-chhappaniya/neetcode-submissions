func findDisappearedNumbers(nums []int) []int {
	mark := make([]bool, len(nums))
	for _, num := range nums {
		// creating boolean mark arr because elements are doubled.
		mark[num-1] = true
	}
	var res []int
	for i:=1; i<= len(nums); i++ {
		if !mark[i-1] {
			// if ele absent then append that index number to result
			res = append(res, i)
		}
	}
	return res
}
