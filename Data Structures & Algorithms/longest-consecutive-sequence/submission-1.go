
func longestConsecutive(nums []int) int {
	hm:=make(map[int]struct{})
	for _, num := range nums {
		hm[num] = struct{}{}
	}
	longest := 0
	for num := range hm {
		if _, ok:=hm[num-1]; !ok {
			length:=1
			for {
				if _, ok:= hm[num+length];ok{
					length++
				} else {
					break
				}
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
