func hasDuplicate(nums []int) bool {
    hm := make(map[int]int)
	for _, v:=range nums{
		if occ,_:=hm[v]; occ >= 1 {
			return true
		}
		hm[v]++
	}
	return false
}
