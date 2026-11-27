func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	i := 2
	for i < num/2 {
		if i*i == num {
			return true
		}
		i += 1
	}
	return false
}
