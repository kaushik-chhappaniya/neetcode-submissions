func kthDistinct(arr []string, k int) string {
	hm := make(map[string]int)
	for _, st := range arr {
	//  create hashmap and if ele not present then only add it.
		hm[st]++
		}
	for _, v:= range arr {
	//  iterate for k times and remove k for present ele. return k == 0 iter valu
		if hm[v] == 1 {
			k--
			if k == 0 {
				return v
			}
		} 
	}
	return ""
}