func topKFrequent(nums []int, k int) []int {
	hm := make(map[int]int)
	// Create hashmap of frequency of each element in arr
	for _,v := range nums {
		hm[v]++
	}
	// Create 2d array containing frequ and value from hashmap 
	//  [[] [1] [2] [3] [] [] []] --> output for given eg input. 
	arr := make([][]int, len(nums)+1)
	for num, freq := range hm {
		arr[freq] = append(arr[freq], num) 
	}
	// Create result arr and add all k elements in the result array. 
	result := []int{}
	for i := len(arr)-1; i >= 0 && len(result) < k; i-- {
		if arr[i] != nil {
			result = append(result, arr[i]...)
		}
	}
	return result
	
}
