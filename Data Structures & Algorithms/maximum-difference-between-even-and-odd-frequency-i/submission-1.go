func maxDifference(s string) int {
	hm := make(map[rune]int)
	for _,v := range s {
		hm[v]++
	}
	var odd, even []int
	for _, v := range hm {
		fmt.Println(v)
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	sort.Ints(odd)
	sort.Ints(even)

return odd[len(odd)-1] - even[0]
}
