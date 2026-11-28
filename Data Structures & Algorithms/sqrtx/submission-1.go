func mySqrt(x int) int {
	// if x < 2 return x
	if x < 2 {
		return x
	}
	// sqrt(x) = 2 * sqrt(x / 4)
	// x >> 2 == dividing by 4
	// << 1 multiplying by 2
	l := mySqrt(x >> 2) << 1
	r := l + 1 //go for next number
	if int64(r) * int64(r) > int64(x) {
		return l
	}
	return r
}
