func minEatingSpeed(piles []int, h int) int {
	// speed := 1
	// for {
	// 	totalTime := 0
	// 	for _, pile := range piles {
	// 		totalTime += int(math.Ceil(float64(pile) / float64(speed)))
	// 	}
	// 	if totalTime <= h {
	// 		return speed
	// 	}
	// 	speed += 1
	// }
	// return speed

	left := 1
	right := 0
	for _, p := range piles {
		if p > right {
			right = p
		}
	} 
	res := right
	for left <= right {
		k := (left + right ) / 2
		totalTime := 0
		for _, p := range piles {
			totalTime += int(math.Ceil(float64(p)/ float64(k)))
		}
		if totalTime <= h {
			res = k
			right = k - 1
		} else {
			left = k + 1
		}
	}
	return res
}
