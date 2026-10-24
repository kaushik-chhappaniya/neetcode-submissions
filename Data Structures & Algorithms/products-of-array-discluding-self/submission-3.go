func productExceptSelf(nums []int) []int {
	pref, suf := make([]int, 0), make([]int, len(nums))
	pref = append(pref, 1)
	suf[len(nums)-1] = 1
	for i, j := 1, len(nums)-2; i < len(nums) && j >=0 ; i,j = i+1, j-1 {
		suf[j]=1
		pref = append(pref, pref[i-1]*nums[i-1])
		suf[j] = suf[j+1] * nums[j+1]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = pref[i] * suf[i]
	}
	return nums
}
