func twoSum(nums []int, target int) []int {
    hm := make(map[int]int)
    for i,v := range(nums) {
        if _, ok := hm[target-v]; ok {
            return []int{hm[target-v], i}
        } 
        hm[v] = i
    }
    return []int{0,0}
}
