func hasDuplicate(nums []int) bool {
    var hm = make(map[int]int)
    for _, v := range(nums) {
        if _, ok := hm[v]; ok {
            return true
        }
        hm[v]+=1
    }
    return false
}
