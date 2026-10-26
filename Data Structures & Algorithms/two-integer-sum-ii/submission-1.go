func twoSum(numbers []int, target int) []int {
	var i,j  int = 0, len(numbers)-1
	for i<=j {
		if numbers[i] + numbers[j] < target {
			i+=1
		} else if numbers[i] + numbers[j] > target {
			j-=1
		} else {
			return []int{i+1, j+1}
		}
	 }
	 return []int{-1,-1}
}
