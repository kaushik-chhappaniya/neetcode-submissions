func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}

	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			stackId := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[stackId] = i - stackId 
		}
		stack = append(stack, i)
	}
	return res
}

	// var arr []int
	// n:=len(temperatures)
	// for curr := 0; curr < n; curr++ {
	// 	count := 1
	// 	j := curr+1
	// 	for j < n {
	// 		if temperatures[j] > temperatures[curr] {
	// 			break
	// 		} 	
	// 		j++
	// 		count++
	// 	}
	// 	if j==n {
	// 		count = 0
	// 	}
	// 	arr = append(arr, count)
	// }
	// return arr