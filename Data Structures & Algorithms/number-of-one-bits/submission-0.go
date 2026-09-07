func hammingWeight(n int) int {
	count := 0
	for {
		if n%2==1{
		count++
		} 
		if n == 0 {
			break
		}
		n/=2
	}
	return count

}
