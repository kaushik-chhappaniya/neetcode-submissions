func isValid(s string) bool {
    if len(s) %2 != 0 {
		return false
	}
	stack := []rune{}
	hm := map[rune]rune{')':'(', ']':'[', '}':'{'}
	for _,c := range s {
		if open, exists := hm[c]; exists {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != open {
					return false
				} 
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
