
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hm := make(map[byte]byte)
	hm2 := make(map[byte]byte)
	for i:=0; i< len(s); i++ {
		c1,c2 :=s[i], t[i]
		if val, ok := hm[c1];ok && val != c2 {
			return false
		}
		if val, ok := hm2[c2];ok && val != c1 {
			return false
		}
		hm[c1] = c2
		hm2[c2] = c1
	}
	return true
}
