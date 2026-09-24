func replaceElements(arr []int) []int {
	ans := make([]int, len(arr))
	right := -1
	for i:=len(arr)-1; i>=0; i-- {
		ans[i] = right
		if right < arr[i] {
			right = arr[i] 
		}
	}
	return ans
}