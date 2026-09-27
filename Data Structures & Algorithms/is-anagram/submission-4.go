func isAnagram(st string, tt string) bool {
	if len(st) != len(tt) {
		return false
	}
	s := strings.ToLower(st)
	t:= strings.ToLower(tt)
	hm:=make(map[rune]int)
	for _, ch := range s {
        hm[ch]++
    }

    for _, ch := range t {
        count, ok := hm[ch]

        if !ok {
            return false
        }

        if count <= 1 {
            delete(hm, ch)
        } else {
            hm[ch]--
        }
    }

    return len(hm) == 0
}