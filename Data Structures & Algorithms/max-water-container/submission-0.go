func maxArea(heights []int) int {
	var maxWater int = 0
	for i,j := 0, len(heights)-1; i<=j; {
		cap := (j-i) * min(heights[i], heights[j])
		maxWater = max(maxWater, cap)
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return maxWater
}
