func asteroidCollision(asteroids []int) []int {
	j := -1
	for _,a := range asteroids {
		for j >=0 && asteroids[j] > 0 && a < 0 {
			if asteroids[j] > -a {
				a = 0
				break
			} else if asteroids[j] == -a {
				j--
				a = 0
				break
			} else {
				j--
			}
		}
		if a != 0 {
			j++
			asteroids[j] = a
		}
	}
	return asteroids[:j+1]
}