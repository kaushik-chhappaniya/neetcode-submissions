func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 1, num
	for l <=r {
		mid := l + (r-l)/2
		sq := mid*mid
		if sq > num {
			r =  mid -1
		} else if sq < num {
			l = mid+1
		} else {
			return true
		}
		}
	return false
}
