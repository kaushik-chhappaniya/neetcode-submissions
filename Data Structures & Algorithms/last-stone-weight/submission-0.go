func lastStoneWeight(stones []int) int {
	maxStone := 0
	for _, stone := range stones {
		maxStone = max(stone, maxStone)
	}
	bucket := make([]int, maxStone+1)
	for _, stone := range stones {
		bucket[stone]++
	}
	first, second := maxStone, maxStone
	for first > 0 {
		if bucket[first]%2 == 0 {
			first --
			continue
		}
		j := min(first-1, second)
		for j > 0 && bucket[j] == 0 {
			j--
		}
		if j== 0 {
			return first
		}
		second = j
		bucket[first]--
		bucket[second]--
		bucket[first-second]++
		first = max(first-second, second)
	}
	return first
}
