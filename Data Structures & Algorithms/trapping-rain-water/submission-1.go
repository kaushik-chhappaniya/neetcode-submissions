func trap(height []int) int {
	n := len(height)
    if n <= 2 {
        return 0
    }
    
	
	prefSum, suffSum, water := make([]int, n), make([]int, n),0
	
	prefSum[0] = height[0]
	suffSum[n-1] = height[n-1]

	for i:=1; i < n; i++ {
		prefSum[i] = max(height[i], prefSum[i-1])
	}
	for i:=n-2; i >= 0;i-- {
		suffSum[i] = max(height[i],suffSum[i+1])
	}
	for i:=0; i < n; i++{
		water += min(prefSum[i], suffSum[i]) - height[i]
	}
	return water

}